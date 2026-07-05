package dto

import (
	"time"
	"tlms/internal/models"

	"github.com/google/uuid"
)

type RegisterToolRequest struct {
	Code            string                      `json:"code" binding:"required"`
	Name            string                      `json:"name" binding:"required"`
	Description     string                      `json:"description" binding:"required"`
	Model           string                      `json:"model" binding:"required"`
	Brand           string                      `json:"brand" binding:"required"`
	Category        models.ToolsCategory        `json:"category" binding:"required"`
	Price           float64                     `json:"price" binding:"required"`
	PhotoID         int64                       `json:"photo_id" binding:"required"`
	UsagePeriod     int64                       `json:"usage_period" binding:"required"`
	UsagePeriodUnit models.ToolsUsagePeriodUnit `json:"usage_period_unit" binding:"required"`
}

type ToolsResponse struct {
	ID              uuid.UUID                   `json:"id"`
	Code            string                      `json:"code"`
	Name            string                      `json:"name"`
	Description     string                      `json:"description"`
	Model           string                      `json:"model"`
	Brand           string                      `json:"brand"`
	Category        models.ToolsCategory        `json:"category"`
	Price           float64                     `json:"price"`
	PhotoID         *int64                      `json:"photo_id"`
	UsagePeriod     int32                       `json:"usage_period"`
	UsagePeriodUnit models.ToolsUsagePeriodUnit `json:"usage_period_unit"`
	CreatedAt       time.Time                   `json:"created_at"`
	CreatedBy       *int64                      `json:"created_by"`
	UpdatedAt       *time.Time                  `json:"updated_at"`
	UpdatedBy       *int64                      `json:"updated_by"`
}
