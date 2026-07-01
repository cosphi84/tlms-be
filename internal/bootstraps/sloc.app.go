package bootstraps

import (
	"tlms/internal/handlers"
	"tlms/internal/repositories"
	"tlms/internal/services"

	"gorm.io/gorm"
)

func InitSlocModule(app *App, db *gorm.DB) {
	rep := repositories.NewStorageLocationRepository(db)
	svc := services.NewStorageLocationService(rep)
	app.StorageLocationHandler = handlers.NewStorageLocationHandler(svc)

}
