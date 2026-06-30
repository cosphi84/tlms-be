package bootstraps

import (
	"tlms/internal/handlers"
	"tlms/internal/repositories"
	"tlms/internal/services"

	"gorm.io/gorm"
)

func InitAuthenticateModule(app *App, db *gorm.DB) {
	userRepos := repositories.NewUserRepository(db)
	authService := services.NewAuthenticateService(app.Authz, userRepos)
	app.AuthenticateHandler = handlers.NewAuthenticateHandler(authService)
}
