package repositories

import (
	"tlms/internal/dto"

	"gorm.io/gorm"
)

func Paginate(value interface{}, pagination *dto.PaginationRequest, db *gorm.DB) (*dto.PaginationResponse, error) {
	var totalRows int64

	if pagination.Page <= 0 {
		pagination.Page = 1
	}

	if pagination.Limit <= 0 {
		pagination.Limit = 10
	}

	if pagination.SortedBy == "" {
		pagination.SortedBy = "created_at"
	}

	if pagination.SortDir == "" {
		pagination.SortDir = "desc"
	}

	offset := (pagination.Page - 1) * pagination.Limit

	if err := db.Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if err := db.Offset(offset).Limit(pagination.Limit).Find(value).Error; err != nil {
		return nil, err
	}

	totalPages := int((totalRows + int64(pagination.Limit) - 1) / int64(pagination.Limit))

	return &dto.PaginationResponse{
		Data:       value,
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		TotalRows:  totalRows,
		TotalPages: totalPages,
	}, nil
}
