package routes

import (
	"tlms/internal/auth"
	"tlms/internal/handlers"
	"tlms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type StorageLocationRouteConfig struct {
	SlocHandler *handlers.StorageLocationHandler
	Authz       *auth.Service
}

func RegisterStorageLocationRoutes(rg *gin.RouterGroup, cfg StorageLocationRouteConfig) {
	slocRouter := rg.Group("/slocs")
	slocRouter.Use(middleware.Authenticate(), middleware.Authorize(cfg.Authz))
	{
		slocRouter.GET("", cfg.SlocHandler.GetAllSloc)
		slocRouter.POST("", cfg.SlocHandler.CreateSloc)
		slocRouter.GET("/:id", cfg.SlocHandler.GetSLoc)
		slocRouter.PUT("/:id", cfg.SlocHandler.UpdateSLoc)
		slocRouter.DELETE("/:id", cfg.SlocHandler.DeleteSLoc)
	}
}
