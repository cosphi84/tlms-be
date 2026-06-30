package routes

import (
	"tlms/internal/auth"
	"tlms/internal/handlers"
	"tlms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type OfficeRouteConfig struct {
	OfficeHandler *handlers.OfficeHandler
	Authz         *auth.Service
}

func RegisterOfficeRoutes(rg *gin.RouterGroup, config OfficeRouteConfig) {
	offices := rg.Group("/offices")
	offices.Use(middleware.Authenticate(), middleware.Authorize(config.Authz))
	{
		offices.POST("", config.OfficeHandler.Create)
		offices.GET("", config.OfficeHandler.FindAll)
		offices.GET("/options", config.OfficeHandler.FindOptions)
		offices.PUT("/:id", config.OfficeHandler.Update)
		offices.DELETE("/:id", config.OfficeHandler.Delete)
	}
}
