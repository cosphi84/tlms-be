package bootstraps

import (
	"tlms/internal/handlers"
	"tlms/internal/repositories"
	"tlms/internal/services"

	"gorm.io/gorm"
)

func InitToolsModule(app *App, db *gorm.DB) {
	rep := repositories.NewToolsRepository(db)
	svc := services.NewToolsService(rep)
	app.ToolsHandler = handlers.NewToolHandler(svc)
}
