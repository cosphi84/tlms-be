package bootstraps

import (
	"tlms/internal/handlers"
	"tlms/internal/repositories"
	"tlms/internal/services"
	"tlms/internal/storage"
	"tlms/internal/validators"

	"gorm.io/gorm"
)

// InitFileModule wires Storage → Validator → Repository → Service → Handler.
// Storage backend dipilih di sini (satu-satunya titik ganti implementasi
// saat migrasi ke MinIO/S3 nanti) — tidak ada bagian lain dari codebase
// yang perlu disentuh.
func InitFileModule(app *App, db *gorm.DB) {
	localStorage, err := storage.NewLocalStorage()
	if err != nil {
		panic(err)
	}

	fileValidator := validators.NewFileValidator()
	fileRepo := repositories.NewFileRepository(db)
	fileSvc := services.NewFileService(fileRepo, localStorage, fileValidator)
	app.FileHandler = handlers.NewFileHandler(fileSvc)
}
