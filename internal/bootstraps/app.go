package bootstraps

import (
	"tlms/internal/handlers"

	"gorm.io/gorm"
)

type App struct {
	AuthenticateHandler *handlers.AuthenticateHandler
	OfficeHandler       *handlers.OfficeHandler
}

func NewApp(db *gorm.DB) *App {
	app := &App{}

	InitAuthenticateModule(app, db)
	InitOfficeModule(app, db)

	return app
}
