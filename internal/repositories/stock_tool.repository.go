package repositories

import (
	"errors"
	"fmt"
	"tlms/internal/dto"
	"tlms/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StokToolRepository interface {
	Create(stockTool *models.StockTools) error
	FindById(id uuid.UUID) (*models.StockTools, error)
	FindAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	FindByEntity(entity *dto.FindStockToolByEntity) (*models.StockTools, error)
	Update(stockTool *models.StockTools) error
	Delete(stockTool *models.StockTools) error
}

type stokToolRepository struct {
	db *gorm.DB
}

func NewStokToolRepository(db *gorm.DB) StokToolRepository {
	return &stokToolRepository{db: db}
}

func (r *stokToolRepository) Create(stockTool *models.StockTools) error {
	return r.db.Create(stockTool).Error
}

func (r *stokToolRepository) FindById(id uuid.UUID) (*models.StockTools, error) {
	var stock models.StockTools
	err := r.db.
		Preload("ToolsData").
		Preload("StorageLocation").
		Where("id = ?", id).First(&stock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &stock, nil
}

func (r *stokToolRepository) FindAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	var stockTools []models.StockTools
	q := r.db.Model(&models.StockTools{}).
		Preload("ToolsData").
		Preload("StorageLocation").
		Where("deleted_at is null").
		Order(fmt.Sprintf("%s %s", pagination.SortedBy, pagination.SortDir))
	return Paginate(&stockTools, pagination, q)
}

func (r *stokToolRepository) FindByEntity(entity *dto.FindStockToolByEntity) (*models.StockTools, error) {
	var stockTool models.StockTools

	tx := r.db.Model(&models.StockTools{}).Where("deleted_at is null")

	if entity.ToolID != uuid.Nil {
		tx = tx.Where("tool_id = ?", entity.ToolID)
	}
	if entity.StorageLocation != 0 {
		tx = tx.Where("storage_location = ?", entity.StorageLocation)
	}

	err := tx.Preload("ToolsData").First(&stockTool).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &stockTool, nil
}

func (r *stokToolRepository) Update(stockTool *models.StockTools) error {
	return r.db.Model(&models.StockTools{}).Updates(stockTool).Error
}

func (r *stokToolRepository) Delete(stockTool *models.StockTools) error {
	return r.db.Delete(stockTool).Error
}
