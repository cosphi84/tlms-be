package bootstraps

import (
	"tlms/internal/handlers"
	"tlms/internal/repositories"
	"tlms/internal/services"

	"gorm.io/gorm"
)

func InitStockToolModule(app *App, db *gorm.DB) {
	stockRepos := repositories.NewStokToolRepository(db)
	toolsRepos := repositories.NewToolsRepository(db)
	slocRepos := repositories.NewStorageLocationRepository(db)

	svc := services.NewStockToolService(stockRepos, toolsRepos, slocRepos)
	app.StockToolHandler = handlers.NewStockToolHandler(svc)
}
