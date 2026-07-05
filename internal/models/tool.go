package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tools struct {
	ID              uuid.UUID            `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Code            string               `gorm:"type:varchar(255);unique;not null;index:idx_tools_code" json:"code"`
	Name            string               `gorm:"type:varchar(255);not null;index:idx_tools_name" json:"name"`
	Description     *string              `gorm:"type:text" json:"description,omitempty"`
	Brand           *string              `gorm:"type:varchar(255)" json:"brand,omitempty"`
	Category        ToolsCategory        `gorm:"type:varchar(100);not null;default:'primary';index:idx_tools_category" json:"category"`
	Price           float64              `gorm:"type:numeric(18,2);not null;default:0" json:"price"`
	UsagePeriod     int32                `gorm:"type:int;not null;default:1" json:"usage_period"`
	UsagePeriodUnit ToolsUsagePeriodUnit `gorm:"type:varchar(255);not null;default:'Y'" json:"usage_period_unit"`
	PhotoID         *int64               `gorm:"column:photo_id;index:idx_tools_photo_id" json:"photo_id,omitempty"`
	PhotoTool       *UploadFile          `gorm:"foreignKey:PhotoID;references:ID;constraint:OnDelete:SET NULL" json:"photo,omitempty"`

	IsActive  bool           `gorm:"not null;default:true;index:idx_tools_is_active" json:"is_active"`
	CreatedAt time.Time      `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"null;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"null;default:null" json:"deleted_at"`
	CreatedBy *int64         `gorm:"column:created_by" json:"created_by,omitempty"`
	UpdatedBy *int64         `gorm:"column:updated_by" json:"updated_by,omitempty"`
	DeletedBy *int64         `gorm:"column:deleted_by" json:"deleted_by,omitempty"`

	CreatedByUser *User `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser *User `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
}

func (Tools) TableName() string {
	return "tools"
}
