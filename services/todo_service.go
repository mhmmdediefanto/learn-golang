package services

import (
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/repositories"
)

type TodoService struct {
	repo repositories.TodoRepository
}

func (s *TodoService) GetAll(userID uint) ([]models.Todo, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoService) Create(todo *models.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) Delete(id uint, userID uint) error {
	return s.repo.Delete(id, userID)
}

func (s *TodoService) Update(id uint, todo *models.Todo, userID uint) (*models.Todo, error) {
	return s.repo.Update(id, todo, userID)
}
