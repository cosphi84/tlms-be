package routes

import (
	"tlms/internal/auth"
	"tlms/internal/handlers"
	"tlms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type FileRouteConfig struct {
	FileHandler *handlers.FileHandler
	Authz       *auth.Service
}

func RegisterFileRoutes(rg *gin.RouterGroup, cfg FileRouteConfig) {
	fileRouter := rg.Group("/files")
	fileRouter.Use(middleware.Authenticate(), middleware.Authorize(cfg.Authz))
	{
		fileRouter.POST("", cfg.FileHandler.Upload)
		fileRouter.GET("/:uuid", cfg.FileHandler.GetMetadata)
		fileRouter.GET("/:uuid/download", cfg.FileHandler.Download)
		fileRouter.DELETE("/:uuid", cfg.FileHandler.Delete)
	}
}
