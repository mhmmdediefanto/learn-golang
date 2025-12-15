package services

import (
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/repositories"
)

type categoryService struct {
	repo repositories.CategoryRepository
}

type CategoryService interface {
	GetAll() ([]models.Category, error)
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}
