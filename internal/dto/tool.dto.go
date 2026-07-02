package dto

import (
	"time"
	"tlms/internal/models"

	"github.com/google/uuid"
)

type RegisterToolRequest struct {
	Code        string               `json:"code" binding:"required"`
	Name        string               `json:"name" binding:"required"`
	Description string               `json:"description" binding:"required"`
	Brand       string               `json:"brand" binding:"required"`
	Category    models.ToolsCategory `json:"category" binding:"required"`
	Price       float64              `json:"price" binding:"required"`
	PhotoID     int64                `json:"photo_id" binding:"required"`
}

type RegisterToolResponse struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	PhotoID     int64     `json:"photo_id"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}
