package models

import (
	"time"
)

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `gorm:"type:varchar(255);unique;not null;index:idx_users_email" json:"email"`
	Password string `gorm:"type:text;not null" json:"-"`

	Name  string  `gorm:"type:varchar(255);not null" json:"name"`
	Image *string `gorm:"type:text" json:"image,omitempty"`

	OfficeID int64 `gorm:"type:integer;not null;index:idx_users_office_id" json:"office_id"`
	IsActive bool  `gorm:"not null;default:true;index:idx_users_is_active" json:"is_active"`

	FailedLoginAttempts int `gorm:"not null;default:0" json:"failed_login_attempts"`

	LockedUntil *time.Time `gorm:"type:timestamp" json:"locked_until,omitempty"`

	LastLoginAt *time.Time `gorm:"type:timestamp" json:"last_login_at,omitempty"`

	LastLoginFrom *string `gorm:"type:varchar" json:"last_login_from,omitempty"`

	CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:now()" json:"updated_at"`

	CreatedBy *int64 `gorm:"column:created_by" json:"created_by,omitempty"`
	UpdatedBy *int64 `gorm:"column:updated_by" json:"updated_by,omitempty"`

	// Relations
	Office Office `gorm:"foreignKey:OfficeID" json:"office"`

	CreatedByUser *User `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser *User `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
}

func (User) TableName() string {
	return "users"
}
