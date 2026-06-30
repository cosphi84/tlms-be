package repositories

import (
	"errors"
	"tlms/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]*models.User, error)
	FindByID(id int32) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id int32) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id int32) (*models.User, error) {
	var user models.User
	err := r.db.
		Preload("Office").
		First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.
		Preload("Office").
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(usr *models.User) error {
	return r.db.Model(usr).
		Select(
			"email",
			"name",
			"password",
			"image",
			"office_id",
			"is_active",
			"failed_login_attempts",
			"locked_until",
			"last_login_at",
			"last_login_from",
			"updated_at",
			"updated_by",
		).Updates(usr).Error
}

func (r *userRepository) Delete(id int32) error {
	return r.db.Delete(&models.User{}, id).Error
}
