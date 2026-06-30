package seeder

import (
	"errors"
	"fmt"
	"os"
	"tlms/internal/auth"
	"tlms/internal/dto"
	"tlms/internal/repositories"
	"tlms/internal/services"

	"gorm.io/gorm"
)

func InitUserAdmin(db *gorm.DB) error {
	usrName := os.Getenv("SEEDER_ADMIN_NAME")
	usrPass := os.Getenv("SEEDER_ADMIN_PASSWORD")
	usrEmail := os.Getenv("SEEDER_ADMIN_EMAIL")
	usrOffice := os.Getenv("SEEDER_OFFICE_HQ_CODE")

	if usrName == "" && usrEmail == "" && usrOffice == "" {
		panic("admin user data empty. set it in the env")
	}

	userRepo := repositories.NewUserRepository(db)
	usrSvc := services.NewUserService(userRepo)

	// roles
	enforcer, err := auth.NewEnforcer(db)
	if err != nil {
		return err
	}
	authSvc := auth.NewService(enforcer)

	usrObj := dto.CreateUserDTO{
		Name:     usrName,
		Email:    usrEmail,
		Password: usrPass,
		OfficeID: 1,
	}

	err = usrSvc.Create(&usrObj)

	if err != nil {
		return err
	}

	savedUser, err := usrSvc.FindByEmail(usrEmail)
	if err != nil {
		return err
	}

	res, err := authSvc.GrantRole(savedUser.Email, string(auth.RoleSuperadmin))
	if err != nil {
		return err
	}
	if !res {
		return errors.New(fmt.Sprintf("user %s failed to grant role", usrName))
	}
	fmt.Println("user successfully granted")
	return nil
}
