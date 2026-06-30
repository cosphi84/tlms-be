package bootstraps

import (
	"tlms/internal/auth"

	"github.com/casbin/casbin/v3"
	"gorm.io/gorm"
)

type SeedApp struct {
	Enforcer *casbin.Enforcer
	Authz    *auth.Service
}

func NewSeedApp(db *gorm.DB) *SeedApp {
	enforcer, err := auth.NewEnforcer(db)
	if err != nil {
		panic(err)
	}
	autz := auth.NewService(enforcer)

	return &SeedApp{
		Enforcer: enforcer,
		Authz:    autz,
	}
}
