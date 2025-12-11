package services

import (
	"errors"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/repositories"
	"go-bakcend-todo-list/utils"

	"gorm.io/gorm"
)

type UserService struct {
	repo repositories.UserRepository
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) Create(user *models.User) (*models.User, error) {

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
