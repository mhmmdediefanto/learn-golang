package services

import (
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/repositories"
)

type todoService struct {
	repo repositories.TodoRepository
}
type TodoService interface {
	GetAll(userID uint) ([]models.Todo, error)
	Create(todo *models.Todo) error
	Delete(id uint, userID uint) error
	Update(id uint, todo *models.Todo, userID uint) (*models.Todo, error)
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) GetAll(userID uint) ([]models.Todo, error) {
	return s.repo.GetAll(userID)
}

func (s *todoService) Create(todo *models.Todo) error {
	return s.repo.Create(todo)
}

func (s *todoService) Delete(id uint, userID uint) error {
	return s.repo.Delete(id, userID)
}

func (s *todoService) Update(id uint, todo *models.Todo, userID uint) (*models.Todo, error) {
	return s.repo.Update(id, todo, userID)
}
