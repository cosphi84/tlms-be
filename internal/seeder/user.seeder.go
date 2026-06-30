package seeder

import (
	"errors"
	"fmt"
	"os"
	"tlms/internal/auth"
	"tlms/internal/bootstraps"
	"tlms/internal/dto"
	"tlms/internal/repositories"
	"tlms/internal/services"

	"gorm.io/gorm"
)

func InitUserAdmin(app *bootstraps.SeedApp, db *gorm.DB) error {
	usrName := os.Getenv("SEEDER_ADMIN_NAME")
	usrPass := os.Getenv("SEEDER_ADMIN_PASSWORD")
	usrEmail := os.Getenv("SEEDER_ADMIN_EMAIL")
	usrOffice := os.Getenv("SEEDER_OFFICE_HQ_CODE")

	if usrName == "" && usrEmail == "" && usrOffice == "" {
		return errors.New("no seed user admin found")
	}

	userRepo := repositories.NewUserRepository(db)
	usrSvc := services.NewUserService(userRepo, app.Authz)

	usrObj := dto.CreateUserDTO{
		Name:     usrName,
		Email:    usrEmail,
		Password: usrPass,
		Role:     string(auth.RoleSuperadmin),
		OfficeID: 1,
	}

	err := usrSvc.Create(&usrObj)

	if err != nil {
		return err
	}

	fmt.Println("user successfully granted")
	return nil
}
