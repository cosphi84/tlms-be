package bootstraps

import (
	"tlms/internal/handlers"
	"tlms/internal/repositories"
	"tlms/internal/services"

	"gorm.io/gorm"
)

func InitUserModule(app *App, db *gorm.DB) {
	userRep := repositories.NewUserRepository(db)
	userSvc := services.NewUserService(userRep, app.Authz)
	app.UserHandler = handlers.NewUserHandler(userSvc)

}
