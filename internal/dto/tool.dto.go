package dto

import (
	"tlms/internal/models"
)

type RegisterToolRequest struct {
	Code            string                      `json:"code" binding:"required"`
	Name            string                      `json:"name" binding:"required"`
	Description     string                      `json:"description" binding:"required"`
	Brand           string                      `json:"brand" binding:"required"`
	Category        models.ToolsCategory        `json:"category" binding:"required"`
	Price           float64                     `json:"price" binding:"required"`
	PhotoID         int64                       `json:"photo_id" binding:"required"`
	UsagePeriod     int64                       `json:"usage_period" binding:"required"`
	UsagePeriodUnit models.ToolsUsagePeriodUnit `json:"usage_period_unit" binding:"required"`
}
