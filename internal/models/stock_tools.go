package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockTools struct {
	ID                uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ToolID            uuid.UUID `json:"tool_id" gorm:"type:uuid;not null;index:idx_stock_tools_id"`
	StorageLocationID int64     `json:"storage_loc_id" gorm:"type:bigint;not null;index:idx_stock_s_loc_id"`

	ToolsData       *Tools           `json:"tools_data,omitempty" gorm:"foreignKey:ToolID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	StorageLocation *StorageLocation `json:"storage_location,omitempty" gorm:"foreignKey:StorageLocationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	Qty  int32  `json:"qty" gorm:"not null;default:0"`
	SBin string `json:"sbin" gorm:"type:varchar(255);not null;default:''"`

	ReferenceType   StockToolsReferenceType `json:"reference_type" gorm:"type:varchar(50);default:'initial_stock'"`
	StockCounter    int32                   `json:"stock_counter" gorm:"not null;default:1"`
	ReferenceNumber *string                 `json:"reference_number" gorm:"type:varchar(255)"`

	CreatedAt time.Time      `json:"created_at" gorm:"not null;default:now()"`
	CreatedBy *int64         `json:"created_by"`
	UpdatedAt *time.Time     `json:"updated_at"`
	UpdatedBy *int64         `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	DeletedBy *int64         `json:"deleted_by"`
}

func (StockTools) TableName() string {
	return "stock_tools"
}
