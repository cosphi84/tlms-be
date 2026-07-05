package services

import (
	"context"
	"errors"
	"time"
	"tlms/internal/auth"
	"tlms/internal/dto"
	"tlms/internal/models"
	"tlms/internal/repositories"

	"github.com/google/uuid"
)

type ToolsService interface {
	Create(req *dto.RegisterToolRequest, ctx context.Context) (*dto.ToolsResponse, error)
	Update(id uuid.UUID, req *dto.RegisterToolRequest, ctx context.Context) error
	FindById(id uuid.UUID) (*models.Tools, error)
	FindAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	Delete(id uuid.UUID) error
}

type toolsService struct {
	toolsRepos repositories.ToolsRepository
}

func NewToolsService(toolRepos repositories.ToolsRepository) ToolsService {
	return &toolsService{toolsRepos: toolRepos}
}

func bindToResponse(m *models.Tools) *dto.ToolsResponse {
	return &dto.ToolsResponse{
		ID:              m.ID,
		Name:            m.Name,
		Code:            m.Code,
		Description:     m.Description,
		Brand:           m.Brand,
		Category:        m.Category,
		PhotoID:         m.PhotoID,
		UsagePeriod:     m.UsagePeriod,
		UsagePeriodUnit: m.UsagePeriodUnit,
		CreatedAt:       m.CreatedAt,
		CreatedBy:       m.CreatedBy,
		UpdatedAt:       m.UpdatedAt,
		UpdatedBy:       m.UpdatedBy,
	}
}

func (s *toolsService) Create(req *dto.RegisterToolRequest, ctx context.Context) (*dto.ToolsResponse, error) {
	usr, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, errors.New("invalid claims")
	}
	existing, err := s.toolsRepos.FindByCode(req.Code)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("code already exists")
	}

	tool := &models.Tools{
		Code:        req.Code,
		Name:        req.Name,
		Description: &req.Description,
		Price:       req.Price,
		PhotoID:     &req.PhotoID,
		Brand:       &req.Brand,
		CreatedBy:   &usr.UserID,
		CreatedAt:   time.Now(),
		IsActive:    true,
	}

	if err := s.toolsRepos.Create(tool); err != nil {
		return nil, err
	}
	return bindToResponse(tool), nil
}
func (s *toolsService) Update(id uuid.UUID, req *dto.RegisterToolRequest, ctx context.Context) error {
	usr, err := auth.GetClaims(ctx)
	if err != nil {
		return errors.New("invalid claims")
	}
	tool, err := s.toolsRepos.FindById(id)
	if err != nil {
		return err
	}
	if tool == nil {
		return errors.New("tool not found")
	}
	now := time.Now()
	tool.Code = req.Code
	tool.Name = req.Name
	tool.Description = &req.Description
	tool.Brand = &req.Brand
	tool.Category = req.Category
	tool.Price = req.Price
	tool.PhotoID = &req.PhotoID
	tool.Brand = &req.Brand
	tool.UpdatedBy = &usr.UserID
	tool.UpdatedAt = &now
	return s.toolsRepos.Update(tool)

}
func (s *toolsService) FindById(id uuid.UUID) (*models.Tools, error) {
	return s.toolsRepos.FindById(id)
}
func (s *toolsService) FindAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	return s.toolsRepos.FindAll(pagination)
}
func (s *toolsService) Delete(id uuid.UUID) error {
	tool, err := s.FindById(id)
	if err != nil {
		return err
	}
	if tool == nil {
		return errors.New("tool not found")
	}
	return s.toolsRepos.Delete(id)
}
