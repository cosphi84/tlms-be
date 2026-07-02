package repositories

import (
	"errors"
	"fmt"
	"tlms/internal/dto"
	"tlms/internal/models"

	"gorm.io/gorm"
)

type ToolsRepository interface {
	Create(tool *models.Tools) error
	Update(data *models.Tools) error
	FindById(id int64) (*models.Tools, error)
	FindByCode(code string) (*models.Tools, error)
	FindAll(request *dto.PaginationRequest) (*dto.PaginationResponse, error)
	Delete(id int64) error
}

type toolsRepository struct {
	db *gorm.DB
}

func NewToolsRepository(db *gorm.DB) ToolsRepository {
	return &toolsRepository{db}
}

func (r *toolsRepository) Create(data *models.Tools) error {
	return r.db.Create(&data).Error
}

func (r *toolsRepository) Update(data *models.Tools) error {
	return r.db.Updates(data).Error
}
func (r *toolsRepository) FindById(id int64) (*models.Tools, error) {
	var tool models.Tools
	err := r.db.Preload("PhotoTool").
		Where("id = ?", id).
		First(&tool).Error
	if err != nil {
		return nil, err
	}
	return &tool, nil
}

func (r *toolsRepository) FindByCode(code string) (*models.Tools, error) {
	var tool models.Tools
	err := r.db.Preload("PhotoTool").
		Where("code = ?", code).
		First(&tool).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tool, nil
}

func (r *toolsRepository) FindAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	var tools []models.Tools
	query := r.db.Model(&models.Tools{}).
		Preload("PhotoTool").
		Where("deleted_at is null").
		Order(fmt.Sprintf("%s %s", pagination.SortedBy, pagination.SortDir))
	return Paginate(&tools, pagination, query)
}
func (r *toolsRepository) Delete(id int64) error {
	return r.db.Delete(&models.Tools{}, "id = ?", id).Error
}
