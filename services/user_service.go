package services

import (
	"errors"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/repositories"
	"go-bakcend-todo-list/utils"

	"gorm.io/gorm"
)

type userService struct {
	repo repositories.UserRepository
}

type UserService interface {
	GetAll() ([]models.User, error)
	Create(user *models.User) (*models.User, error)
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) Create(user *models.User) (*models.User, error) {

	existingEmail, err := s.repo.FindByEmail(user.Email)
	// jika error bukan "not found", berarti error database
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingEmail != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	bytePassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = bytePassword
	return s.repo.Create(user)
}
