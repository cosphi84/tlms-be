package auth

import "github.com/golang-jwt/jwt/v5"

type RoleType string
type Key string

const (
	RoleSuperadmin RoleType = "superadmin"
	RoleAdminHQ    RoleType = "admin_hq"
	RoleSvcHead    RoleType = "service_head"
	RoleManagement RoleType = "management"
	RoleAuditor    RoleType = "auditor"
	RoleAuditorCS  RoleType = "auditor_cs"
	RoleTechnician RoleType = "technician"
)

type JWTClaims struct {
	UserID    int64      `json:"user_id"`
	Role      []RoleType `json:"role"`
	OfficeID  int64      `json:"office_id"`
	TokenType string     `json:"token_type"`

	jwt.RegisteredClaims
}

const (
	AuthContextKey Key = "auth"
)
