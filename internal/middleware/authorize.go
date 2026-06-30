package middleware

import (
	"net/http"
	"tlms/internal/auth"

	"github.com/gin-gonic/gin"
)

func Authorize(authz *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		v := c.Request.Context().Value(auth.AuthContextKey)
		if v == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		claims, ok := v.(*auth.JWTClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		obj := c.FullPath()
		act := c.Request.Method

		allowed := false
		for _, role := range claims.Role {
			ok, err := authz.Enforce(string(role), obj, act)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			if ok {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
