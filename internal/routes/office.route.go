package routes

import (
	"tlms/internal/auth"
	"tlms/internal/handlers"
	"tlms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type OfficeRouteConfig struct {
	officeHandler *handlers.OfficeHandler
}

func RegisterOfficeRoutes(rg *gin.RouterGroup, config OfficeRouteConfig) {
	offices := rg.Group("/offices")
	offices.Use(middleware.Authenticate(), middleware.Authorize(&auth.Service{}))
	{
		offices.POST("", config.officeHandler.Create)
		offices.GET("", config.officeHandler.FindAll)
		offices.GET("/options", config.officeHandler.FindOptions)
		offices.PUT("/:id", config.officeHandler.Update)
		offices.DELETE("/:id", config.officeHandler.Delete)
	}
}
