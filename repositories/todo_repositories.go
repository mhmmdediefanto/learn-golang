package repositories

import (
	"errors"
	"go-bakcend-todo-list/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAll(userID uint) ([]models.Todo, error)
	Create(todo *models.Todo) error
	Delete(id uint, userID uint) error
	Update(id uint, data *models.Todo, userID uint) (*models.Todo, error)
}
type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) GetAll(userID uint) ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.
		Where("user_id = ?", userID).
		Preload("User").
		Find(&todos)
	return todos, result.Error
}

func (r *todoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) Delete(id uint, userID uint) error {
	result := r.db.
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&models.Todo{}, id)

	if result.RowsAffected == 0 {
		return errors.New("todo tidak di temukan")
	}

	return result.Error
}

func (r *todoRepository) Update(id uint, data *models.Todo, userID uint) (*models.Todo, error) {
	result := r.db.Model(&models.Todo{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(data)

	if result.RowsAffected == 0 {
		return nil, errors.New("todo tidak ditemukan")
	}

	// Ambil data terbaru
	var updatedTodo models.Todo
	if err := r.db.First(&updatedTodo, id).Error; err != nil {
		return nil, err
	}

	return &updatedTodo, nil
}
