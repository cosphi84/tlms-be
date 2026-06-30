package bootstraps

import (
	"tlms/internal/handlers"
	"tlms/internal/repositories"
	"tlms/internal/services"

	"gorm.io/gorm"
)

// InitOfficeModule wires the repository → service → handler dependency chain
// and returns a ready-to-use OfficeHandler.
//
// Call this from app.go during application startup.
func InitOfficeModule(app *App, db *gorm.DB) {
	repo := repositories.NewOfficeRepository(db)
	svc := services.NewOfficeService(repo)
	app.OfficeHandler = handlers.NewOfficeHandler(svc)
}
