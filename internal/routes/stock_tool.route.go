package routes

import (
	"tlms/internal/auth"
	"tlms/internal/handlers"
	"tlms/internal/middleware"

	"github.com/gin-gonic/gin"
)

type StockToolsRouteConfig struct {
	StockToolHandler *handlers.StockToolHandler
	Authz            *auth.Service
}

func RegisterStockToolRoutes(r *gin.RouterGroup, cfg StockToolsRouteConfig) {
	stockTools := r.Group("/stock-tools")
	stockTools.Use(middleware.Authenticate(), middleware.Authorize(cfg.Authz))
	{
		stockTools.POST("/add", cfg.StockToolHandler.AddStockTool)
		stockTools.POST("/remove", cfg.StockToolHandler.RemoveStockTool)
	}
}
