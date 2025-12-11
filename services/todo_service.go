package services

import (
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/repositories"
)

type TodoService struct {
	repo repositories.TodoRepository
}

func (s *TodoService) GetAll() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) Create(todo *models.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *TodoService) Update(id uint, todo *models.Todo) (*models.Todo, error) {
	return s.repo.Update(id, todo)
}
