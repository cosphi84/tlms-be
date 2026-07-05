package repositories

import (
	"errors"
	"fmt"
	"tlms/internal/dto"
	"tlms/internal/models"

	nestedset "github.com/longbridgeapp/nested-set"
	"gorm.io/gorm"
)

type OfficeRepository interface {
	CreateRoot(office *models.Office) error
	CreateChild(office *models.Office, parentOffice *models.Office) error
	FindById(id int64) (*models.Office, error)
	FindByCode(code string) (*models.Office, error)
	FindAll() ([]models.Office, error)
	//FindAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	FindOffices() ([]dto.OfficeOptionResponse, error)
	Update(office *models.Office) error
	Delete(id int64) error
}

type officeRepository struct {
	db *gorm.DB
}

func NewOfficeRepository(db *gorm.DB) OfficeRepository {
	return &officeRepository{db: db}
}

func (repos *officeRepository) CreateRoot(office *models.Office) error {
	tx := repos.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := nestedset.Create(tx, office, nil)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

func (repos *officeRepository) CreateChild(office *models.Office, parentOffice *models.Office) error {
	tx := repos.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := nestedset.Create(tx, office, parentOffice)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (repos *officeRepository) FindById(id int64) (*models.Office, error) {
	var office models.Office

	err := repos.db.
		Where("id=?", id).
		First(&office).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &office, nil
}

func (repos *officeRepository) FindByCode(code string) (*models.Office, error) {
	var office models.Office

	err := repos.db.
		Where("code=?", code).
		First(&office).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &office, nil
}

/*
func (repos *officeRepository) FindAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	var offices []models.Office

	baseQuery := repos.db.Model(&models.Office{}).
		Where("deleted_at IS NULL").
		Order(fmt.Sprintf("%s %s", pagination.SortedBy, pagination.SortDir))

	return Paginate(&offices, pagination, baseQuery)
}
*/

func (repos *officeRepository) FindAll() ([]models.Office, error) {
	var offices []models.Office

	err := repos.db.Model(&models.Office{}).
		Where("deleted_at IS NULL").
		Find(&offices).Error
	if err != nil {
		return nil, err
	}
	return offices, nil
}

func (repos *officeRepository) FindOffices() ([]dto.OfficeOptionResponse, error) {
	type raw struct {
		ID   int64
		Name string
		Code string
	}

	var rows []raw
	err := repos.db.Model(&models.Office{}).
		Where("deleted_at IS NULL").
		Select("id, name, code").
		Order("id ASC").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	opts := make([]dto.OfficeOptionResponse, 0, len(rows))
	for _, row := range rows {
		opts = append(opts, dto.OfficeOptionResponse{
			ID:    row.ID,
			Label: fmt.Sprintf("%s - %s", row.Name, row.Code),
		})
	}
	return opts, nil
}

func (repos *officeRepository) Update(office *models.Office) error {
	return repos.db.Model(office).
		Select("code", "name", "type").
		Updates(office).Error
}

func (repos *officeRepository) Delete(id int64) error {
	return repos.db.Delete(&models.Office{}, id).Error
}
