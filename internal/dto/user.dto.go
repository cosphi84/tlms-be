package dto

type CreateUserDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Image    string `json:"image"`
	OfficeID int64  `json:"office_id"`
	Role     string `json:"role"`
}
