package seeder

import (
	"tlms/internal/repositories"
	"tlms/internal/services"

	"gorm.io/gorm"
)

func InitOfficeSeed(db *gorm.DB) error {
	officeRepo := repositories.NewOfficeRepository(db)
	officeSvc := services.NewOfficeService(officeRepo)

	return officeSvc.CreateHQ()
}
