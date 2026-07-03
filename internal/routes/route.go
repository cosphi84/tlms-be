package routes

import (
	"net/http"
	"os"
	"tlms/internal/bootstraps"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, app *bootstraps.App) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "X-Request-ID"},
	}))

	appVer := os.Getenv("APP_VERSION")
	if appVer == "" {
		appVer = "1.0.0"
	}
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "tms"
	}

	api := router.Group("/")

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to " + appName + " Backend API",
			"version": appVer,
		})
	})

	// Authenticate
	AuthenticateRoute(api, AuthenticateRouteConfig{
		authHandler: app.AuthenticateHandler,
	})

	// Office Management
	RegisterOfficeRoutes(api, OfficeRouteConfig{
		OfficeHandler: app.OfficeHandler,
		Authz:         app.Authz,
	})

	// Users Management
	RegisterUserRoutes(api, UserRouteConfig{
		UserHandler: app.UserHandler,
		Autz:        app.Authz,
	})

	// Storage Loc
	RegisterStorageLocationRoutes(api, StorageLocationRouteConfig{
		SlocHandler: app.StorageLocationHandler,
		Authz:       app.Authz,
	})

	// File Manager
	RegisterFileRoutes(api, FileRouteConfig{
		FileHandler: app.FileHandler,
		Authz:       app.Authz,
	})

	// Tools Manager
	RegisterToolRoutes(api, ToolRouteConfig{
		ToolHandler: app.ToolsHandler,
		Authz:       app.Authz,
	})

	// stock manager
	RegisterStockToolRoutes(api, StockToolsRouteConfig{
		StockToolHandler: app.StockToolHandler,
		Authz:            app.Authz,
	})
}
