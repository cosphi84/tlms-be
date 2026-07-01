package services

import (
	"context"
	"errors"
	"time"
	"tlms/internal/auth"
	"tlms/internal/dto"
	"tlms/internal/models"
	"tlms/internal/repositories"
)

type StogareLocationService interface {
	Create(req dto.StorageLocationRequest, ctx context.Context) error
	FindById(id int64) (*models.StorageLocation, error)
	GetAllSloc(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	Update(id int64, req dto.StorageLocationRequest, ctx context.Context) error
	Delete(id int64, ctx context.Context) error
}

type storageLocationRepository struct {
	slocRepository repositories.StorageLocationRepository
}

func NewStorageLocationService(slocRepository repositories.StorageLocationRepository) StogareLocationService {
	return &storageLocationRepository{slocRepository}
}

func (srv *storageLocationRepository) Create(req dto.StorageLocationRequest, ctx context.Context) error {
	usr, err := auth.GetClaims(ctx)
	if err != nil {
		return errors.New("invalid claims")
	}

	existing, err := srv.slocRepository.FindByCode(req.Code)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("storage location already exists")
	}

	sloc := models.StorageLocation{
		Code:      req.Code,
		Name:      req.Name,
		OfficeID:  req.OfficeID,
		CreatedAt: time.Now(),
		CreatedBy: &usr.UserID,
		IsActive:  true,
	}

	return srv.slocRepository.Create(&sloc)
}

func (srv *storageLocationRepository) FindById(id int64) (*models.StorageLocation, error) {
	return srv.slocRepository.FindById(id)
}

func (srv *storageLocationRepository) GetAllSloc(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	return srv.slocRepository.FindAll(pagination)
}

func (srv *storageLocationRepository) Update(id int64, req dto.StorageLocationRequest, ctx context.Context) error {
	sloc, err := srv.slocRepository.FindById(id)
	if err != nil {
		return err
	}
	if sloc == nil {
		return errors.New("storage location not found")
	}

	if req.Code != "" {
		existing, err := srv.slocRepository.FindByCode(req.Code)
		if err != nil {
			return err
		}
		if existing != nil && existing.ID != id {
			return errors.New("storage location already exists")

		}
		sloc.Code = req.Code
	}
	if req.Name != "" {
		sloc.Name = req.Name
	}

	if req.OfficeID != 0 {
		sloc.OfficeID = req.OfficeID
	}
	return srv.slocRepository.Update(sloc)
}

func (srv *storageLocationRepository) Delete(id int64, ctx context.Context) error {
	sloc, err := srv.slocRepository.FindById(id)
	if err != nil {
		return err
	}
	if sloc == nil {
		return errors.New("storage location not found")
	}

	return srv.slocRepository.Delete(id)
}
