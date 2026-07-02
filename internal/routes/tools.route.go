package routes

import (
	"tlms/internal/auth"
	"tlms/internal/handlers"
	"tlms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type ToolRouteConfig struct {
	ToolHandler *handlers.ToolsHandler
	Authz       *auth.Service
}

func RegisterToolRoutes(router *gin.RouterGroup, cfg ToolRouteConfig) {
	toolRoutes := router.Group("/tools")
	toolRoutes.Use(middleware.Authenticate(), middleware.Authorize(cfg.Authz))

	{
		toolRoutes.GET("", cfg.ToolHandler.GetAllTools)
		toolRoutes.POST("", cfg.ToolHandler.RegisterTool)
		toolRoutes.GET("/:id", cfg.ToolHandler.GetTool)
		toolRoutes.PUT("/:id", cfg.ToolHandler.UpdateTool)
		toolRoutes.DELETE("/:id", cfg.ToolHandler.DeleteTool)
	}
}
