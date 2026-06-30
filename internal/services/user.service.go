package services

import (
	"errors"
	"tlms/internal/auth"
	"tlms/internal/dto"
	"tlms/internal/models"
	"tlms/internal/repositories"
)

type UserService interface {
	Create(usr *dto.CreateUserDTO) error
	FindByEmail(email string) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Create(usr *dto.CreateUserDTO) error {
	existing, err := s.userRepo.FindByEmail(usr.Email)
	if err != nil {
		return err
	}

	if existing != nil {
		return errors.New("user with this email already exists")
	}

	hashedPwd, err := auth.HashPassword(usr.Password)
	if err != nil {
		return err
	}
	user := models.User{
		Email:    usr.Email,
		Name:     usr.Name,
		Password: hashedPwd,
		OfficeID: usr.OfficeID,
		IsActive: true,
		Image:    &usr.Image,
	}

	return s.userRepo.Create(&user)
}

func (s *userService) FindByEmail(email string) (*models.User, error) {
	return s.userRepo.FindByEmail(email)
}
