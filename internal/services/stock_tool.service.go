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

type StockToolService interface {
	ModifyStock(req *dto.ModifyStokToolRequest, ctx context.Context) error
}

type stockToolRepos struct {
	StockRepos      repositories.StokToolRepository
	ToolRepos       repositories.ToolsRepository
	StorageLocRepos repositories.StorageLocationRepository
}

func NewStockToolService(
	stockRepos repositories.StokToolRepository,
	toolRepos repositories.ToolsRepository,
	storageLocRepos repositories.StorageLocationRepository,
) StockToolService {
	return &stockToolRepos{
		StockRepos:      stockRepos,
		ToolRepos:       toolRepos,
		StorageLocRepos: storageLocRepos,
	}
}

func (r *stockToolRepos) ModifyStock(req *dto.ModifyStokToolRequest, ctx context.Context) error {
	usr, err := auth.GetClaims(ctx)
	if err != nil {
		return errors.New("invalid claims")
	}

	theTool, err := r.ToolRepos.FindById(req.ToolID)
	if err != nil {
		return err
	}
	if theTool == nil {
		return errors.New("tool not found")
	}

	theSloc, err := r.StorageLocRepos.FindById(req.SlocID)
	if err != nil {
		return err
	}
	if theSloc == nil {
		return errors.New("storage location not found")
	}

	entity := dto.FindStockToolByEntity{
		StorageLocation: req.SlocID,
		ToolID:          req.ToolID,
	}

	now := time.Now()

	stock, err := r.StockRepos.FindByEntity(&entity)
	if err != nil {
		return err
	}

	if stock == nil {
		dataStock := models.StockTools{
			ToolsData:       theTool,
			StorageLocation: theSloc,
			StockCounter:    1,
			Qty:             req.Qty,
			ReferenceType:   models.RefInitialStock,
			CreatedAt:       now,
			CreatedBy:       &usr.UserID,
		}
		return r.StockRepos.Create(&dataStock)
	}
	if req.Mode == string(models.ModifyStockToolsTypeOutgoing) {
		if stock.Qty > 0 {
			stock.Qty = stock.Qty - req.Qty
		} else {
			stock.Qty = 0
		}

	} else {
		stock.Qty = stock.Qty + req.Qty
	}

	stock.StockCounter = stock.StockCounter + 1
	stock.UpdatedBy = &usr.UserID
	stock.UpdatedAt = &now
	return r.StockRepos.Update(stock)
}
