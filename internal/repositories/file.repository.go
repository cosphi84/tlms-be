package repositories

import (
	"errors"
	"tlms/internal/models"

	"gorm.io/gorm"
)

type FileRepository interface {
	Create(file *models.UploadFile) error
	FindByUUID(uuid string) (*models.UploadFile, error)
	Update(file *models.UploadFile) error
	SoftDelete(id int64, deletedBy *int64) error

	// WithTx mengembalikan instance FileRepository baru yang beroperasi
	// di dalam transaction milik caller (mis. ToolsService.ReplaceIcon).
	// Ini memungkinkan File Manager Module ikut serta dalam atomic
	// transaction lintas-modul tanpa File Manager perlu tahu apa-apa
	// soal domain caller — caller yang mengontrol Begin/Commit/Rollback.
	WithTx(tx *gorm.DB) FileRepository
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db}
}

func (r *fileRepository) WithTx(tx *gorm.DB) FileRepository {
	return &fileRepository{db: tx}
}

func (r *fileRepository) Create(file *models.UploadFile) error {
	return r.db.Create(file).Error
}

func (r *fileRepository) FindByUUID(uuid string) (*models.UploadFile, error) {
	var file models.UploadFile
	err := r.db.Where("uuid = ?", uuid).First(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &file, nil
}

func (r *fileRepository) Update(file *models.UploadFile) error {
	return r.db.Model(file).Updates(file).Error
}

func (r *fileRepository) SoftDelete(id int64, deletedBy *int64) error {
	return r.db.Model(&models.UploadFile{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_by": deletedBy,
			"deleted_at": gorm.Expr("NOW()"),
		}).Error
}
