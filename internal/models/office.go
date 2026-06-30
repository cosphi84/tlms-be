package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Office struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id" nestedset:"id"`

	ParentID sql.NullInt64 `gorm:"column:parent_id" json:"parent_id" nestedset:"parent_id"`

	Code string `gorm:"type:varchar(10);unique;not null;index:idx_offices_code" json:"code"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Type string `gorm:"type:varchar(50);not null" json:"type"`

	Depth int `gorm:"column:depth" json:"depth" nestedset:"depth"`
	Rgt   int `gorm:"column:rgt" json:"rgt" nestedset:"rgt"`
	Lft   int `gorm:"column:lft" json:"lft" nestedset:"lft"`

	ChildrenCount int `gorm:"column:children_count" json:"children_count" nestedset:"children_count"`

	CreatedAt time.Time      `gorm:"not null;default:now()" json:"created_at"`
	CreatedBy *int64         `gorm:"column:created_by" json:"created_by,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	DeletedBy     *string `gorm:"type:varchar" json:"deleted_by,omitempty"`
	CreatedByUser *User   `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	Parent        *Office `gorm:"foreignKey:ParentID" json:"parent,omitempty"`

	Children []Office `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

func (Office) TableName() string {
	return "offices"
}
