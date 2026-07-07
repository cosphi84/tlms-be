package repositories

import (
	"errors"
	"fmt"
	"tlms/internal/dto"
	"tlms/internal/models"

	"gorm.io/gorm"
)

type StorageLocationRepository interface {
	Create(sloc *models.StorageLocation) error
	FindById(id int64) (*models.StorageLocation, error)
	FindByCode(code string) (*models.StorageLocation, error)
	FindAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	Update(sloc *models.StorageLocation) error
	Delete(id int64) error
}

type slocRepository struct {
	db *gorm.DB
}

func NewStorageLocationRepository(db *gorm.DB) StorageLocationRepository {
	return &slocRepository{db}
}

func (r *slocRepository) Create(sloc *models.StorageLocation) error {
	return r.db.Create(sloc).Error
}

func (r *slocRepository) FindById(id int64) (*models.StorageLocation, error) {
	var sloc models.StorageLocation
	err := r.db.
		Preload("Office").
		Where("id=?", id).First(&sloc).Error
	if err != nil {
		return nil, err
	}
	return &sloc, nil
}

func (r *slocRepository) FindByCode(code string) (*models.StorageLocation, error) {
	var loc models.StorageLocation
	err := r.db.Preload("Office").Where("code=?", code).First(&loc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &loc, nil
}

func (r *slocRepository) FindAll(paginate *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	var slocs []models.StorageLocation
	query := r.db.Model(&models.StorageLocation{}).
		Preload("Office").
		Where("deleted_at IS NULL").
		Order(fmt.Sprintf("%s %s", paginate.SortedBy, paginate.SortDir))
	return Paginate(&slocs, paginate, query)
}

func (r *slocRepository) Update(sloc *models.StorageLocation) error {
	return r.db.Model(sloc).
		Updates(sloc).Error
}

func (r *slocRepository) Delete(id int64) error { return r.Delete(id) }
