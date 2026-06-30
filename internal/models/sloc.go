package models

import (
	"time"

	"gorm.io/gorm"
)

type StorageLocation struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	Code     string `gorm:"type:varchar(255);unique;not null;index:idx_storage_locs_code" json:"code"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	OfficeID int64  `gorm:"type:smallint;not null;index:idx_storage_locs_office_id" json:"office_id"`
	IsActive bool   `gorm:"not null;default:true" json:"is_active"`

	CreatedAt time.Time      `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"not null;default:null" json:"deleted_at"`
	CreatedBy *int64         `gorm:"column:created_by" json:"created_by,omitempty"`
	UpdatedBy *int64         `gorm:"column:updated_by" json:"updated_by,omitempty"`
	DeletedBy *int64         `gorm:"column:deleted_by" json:"deleted_by,omitempty"`

	// Relations
	Office        Office `gorm:"foreignKey:OfficeID" json:"office"`
	CreatedByUser *User  `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser *User  `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
	DeletedByUser *User  `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
}

func (StorageLocation) TableName() string {
	return "storage_locations"
}
