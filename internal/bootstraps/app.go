package bootstraps

import (
	"tlms/internal/auth"
	"tlms/internal/handlers"

	"github.com/casbin/casbin/v3"
	"gorm.io/gorm"
)

type App struct {
	Enforcer               *casbin.Enforcer
	Authz                  *auth.Service
	AuthenticateHandler    *handlers.AuthenticateHandler
	OfficeHandler          *handlers.OfficeHandler
	UserHandler            *handlers.UserHandler
	StorageLocationHandler *handlers.StorageLocationHandler
	FileHandler            *handlers.FileHandler
}

func NewApp(db *gorm.DB) *App {
	app := &App{}
	enforcer, err := auth.NewEnforcer(db)
	if err != nil {
		panic(err)
	}
	app.Enforcer = enforcer
	app.Authz = auth.NewService(enforcer)

	InitAuthenticateModule(app, db)
	InitOfficeModule(app, db)
	InitUserModule(app, db)
	InitSlocModule(app, db)
	InitFileModule(app, db)

	return app
}
