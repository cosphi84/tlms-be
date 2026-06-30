package seeder

import (
	"tlms/internal/auth"

	"gorm.io/gorm"
)

func InitCasbinSeed(db *gorm.DB) error {
	enforcer, err := auth.NewEnforcer(db)
	if err != nil {
		return err
	}

	authService := auth.NewService(enforcer)

	// Admin Roles
	adminRoles := []auth.RoleType{
		auth.RoleSuperadmin,
		auth.RoleAdminHQ,
	}

	// Read-only Roles
	readRoles := []auth.RoleType{
		auth.RoleSvcHead,
		auth.RoleManagement,
		auth.RoleAuditor,
		auth.RoleAuditorCS,
		auth.RoleTechnician,
	}

	for _, role := range adminRoles {
		authService.GrantPermission(string(role), "/offices", "GET")
		authService.GrantPermission(string(role), "/offices/options", "GET")
		authService.GrantPermission(string(role), "/offices", "POST")
		authService.GrantPermission(string(role), "/offices/:id", "PUT")
		authService.GrantPermission(string(role), "/offices/:id", "DELETE")
	}

	for _, role := range readRoles {
		authService.GrantPermission(string(role), "/offices", "GET")
		authService.GrantPermission(string(role), "/offices/options", "GET")
	}

	return nil
}
