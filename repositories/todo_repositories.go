package repositories

import (
	"errors"
	"go-bakcend-todo-list/config"
	"go-bakcend-todo-list/models"
)

type TodoRepository struct{}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := config.DB.Find(&todos)
	return todos, result.Error
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return config.DB.Create(todo).Error
}

func (r *TodoRepository) Delete(id uint) error {
	result := config.DB.Delete(&models.Todo{}, id)

	if result.RowsAffected == 0 {
		return errors.New("todo tidak di temukan")
	}

	return result.Error
}

func (r *TodoRepository) Update(id uint, data *models.Todo) (*models.Todo, error) {
	result := config.DB.Model(&models.Todo{}).
		Where("id = ?", id).
		Updates(data)

	if result.RowsAffected == 0 {
		return nil, errors.New("todo tidak ditemukan")
	}

	// Ambil data terbaru
	var updatedTodo models.Todo
	if err := config.DB.First(&updatedTodo, id).Error; err != nil {
		return nil, err
	}

	return &updatedTodo, nil
}
