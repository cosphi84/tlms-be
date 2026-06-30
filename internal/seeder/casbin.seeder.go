package seeder

import (
	"tlms/internal/auth"

	"gorm.io/gorm"
)

func InitCasbinSeed(autz *auth.Service, db *gorm.DB) error {

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
		autz.GrantPermission(string(role), "/offices", "GET")
		autz.GrantPermission(string(role), "/offices/options", "GET")
		autz.GrantPermission(string(role), "/offices", "POST")
		autz.GrantPermission(string(role), "/offices/:id", "PUT")
		autz.GrantPermission(string(role), "/offices/:id", "GET")
		autz.GrantPermission(string(role), "/offices/:id", "DELETE")

		// users
		autz.GrantPermission(string(role), "/users", "POST")
	}

	for _, role := range readRoles {
		autz.GrantPermission(string(role), "/offices", "GET")
		autz.GrantPermission(string(role), "/offices/options", "GET")
		autz.GrantPermission(string(role), "/offices/:id", "GET")
	}

	return nil
}
