package services

import (
	"context"
	"errors"
	"os"
	"time"
	"tlms/internal/auth"
	"tlms/internal/dto"
	"tlms/internal/models"
	"tlms/internal/repositories"
)

type OfficeService interface {
	CreateHQ() error
	CreateOffice(req dto.CreateOfficeRequest, ctx context.Context) error
	GetOffices() ([]models.Office, error)
	//GetOffices(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	GetOfficeOptions() ([]dto.OfficeOptionResponse, error)
	GetOffice(id int64) (*models.Office, error)
	UpdateOffice(id int64, req dto.UpdateOfficeRequest, ctx context.Context) error
	DeleteOffice(id int64, ctx context.Context) error
}

type officeService struct {
	officeRepo repositories.OfficeRepository
}

func NewOfficeService(officeRepo repositories.OfficeRepository) OfficeService {
	return &officeService{officeRepo: officeRepo}
}

func (s *officeService) CreateHQ() error {
	hqName := os.Getenv("SEEDER_OFFICE_HQ_NAME")
	hqCode := os.Getenv("SEEDER_OFFICE_HQ_CODE")
	if hqName == "" || hqCode == "" {
		return errors.New("SEEDER_OFFICE_HQ_NAME and SEEDER_OFFICE_HQ_CODE must be set")
	}

	existing, err := s.officeRepo.FindByCode(hqCode)
	if err != nil {
		return err
	}
	if existing != nil {
		return nil // idempotent — HQ already seeded
	}

	hq := models.Office{
		Name: hqName,
		Code: hqCode,
		Type: "hq",
	}
	return s.officeRepo.CreateRoot(&hq)
}

func (s *officeService) CreateOffice(req dto.CreateOfficeRequest, ctx context.Context) error {
	usr, err := auth.GetClaims(ctx)
	if err != nil {
		return errors.New("auth claims error")
	}

	existing, err := s.officeRepo.FindByCode(req.Code)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("office is exists")
	}

	parent, err := s.officeRepo.FindById(req.ParentID)
	if err != nil {
		return err
	}
	if parent == nil {
		return errors.New("parent office not found")
	}

	newOffice := models.Office{
		Name:      req.Name,
		Code:      req.Code,
		Type:      req.Type,
		Parent:    parent,
		CreatedAt: time.Now(),
		CreatedBy: &usr.UserID,
	}

	return s.officeRepo.CreateChild(&newOffice, parent)
}

func (s *officeService) GetOffices() ([]models.Office, error) {
	return s.officeRepo.FindAll()
}

func (s *officeService) GetOfficeOptions() ([]dto.OfficeOptionResponse, error) {
	return s.officeRepo.FindOffices()
}

func (s *officeService) GetOffice(id int64) (*models.Office, error) {
	return s.officeRepo.FindById(id)
}

func (s *officeService) UpdateOffice(id int64, req dto.UpdateOfficeRequest, ctx context.Context) error {
	office, err := s.officeRepo.FindById(id)
	if err != nil {
		return err
	}
	if office == nil {
		return errors.New("office not found")
	}

	// Apply patch — only update fields that were actually sent in the request body.
	// Empty string check is fine here because code/name/type should never be blank.
	if req.Code != "" {
		// Prevent duplicate code collision on update
		existing, err := s.officeRepo.FindByCode(req.Code)
		if err != nil {
			return err
		}
		if existing != nil && existing.ID != id {
			return errors.New("office with that code already exists")
		}
		office.Code = req.Code
	}
	if req.Name != "" {
		office.Name = req.Name
	}
	if req.Type != "" {
		office.Type = req.Type
	}

	return s.officeRepo.Update(office)
}

func (s *officeService) DeleteOffice(id int64, ctx context.Context) error {
	office, err := s.officeRepo.FindById(id)
	if err != nil {
		return err
	}
	if office == nil {
		return errors.New("office not found")
	}

	// Guard: refuse to delete a node that still has active children.
	// The nested-set lft/rgt math: children exist when rgt - lft > 1.
	if office.Rgt-office.Lft > 1 {
		return errors.New("cannot delete office that still has children")
	}

	return s.officeRepo.Delete(id)
}
