package routes

import (
	"tlms/internal/handlers"

	"github.com/gin-gonic/gin"
)

type AuthenticateRouteConfig struct {
	authHandler *handlers.AuthenticateHandler
}

func AuthenticateRoute(r *gin.RouterGroup, config AuthenticateRouteConfig) {
	authGroup := r.Group("/")

	{
		authGroup.POST("/auth", config.authHandler.Authenticate)
		authGroup.POST("/auth/refresh", config.authHandler.RefreshToken)
	}

}
