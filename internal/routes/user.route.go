package routes

import (
	"tlms/internal/auth"
	"tlms/internal/handlers"
	"tlms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouteConfig struct {
	UserHandler *handlers.UserHandler
	Autz        *auth.Service
}

func RegisterUserRoutes(router *gin.RouterGroup, config UserRouteConfig) {
	usr := router.Group("/users")
	usr.Use(middleware.Authenticate(), middleware.Authorize(config.Autz))

	{
		usr.POST("", config.UserHandler.Create)
	}
}
