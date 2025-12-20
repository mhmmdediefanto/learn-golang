package repositories

import (
	"errors"
	"go-bakcend-todo-list/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	// Define methods for category repository here
	GetAll() ([]models.Category, error)
	Create(category *models.Category) (*models.Category, error)
	Delete(id uint) error
	Update(category *models.Category, id uint) (*models.Category, error)
	FindLatestSlug(slug string) (string, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Create(category *models.Category) (*models.Category, error) {
	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Category{}, id)

	if result.RowsAffected == 0 {
		return errors.New("kategori tidak ditemukan")
	}
	return result.Error
}

func (r *categoryRepository) Update(category *models.Category, id uint) (*models.Category, error) {
	result := r.db.Model(&models.Category{}).
		Where("id = ?", id).
		Updates(category) // GORM otomatis update field non-zero,

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("kategori dengan id tidak ditemukan")
	}

	return category, nil
}

func (r *categoryRepository) FindLatestSlug(baseSlug string) (string, error) {
	var slug string

	err := r.db.Model(&models.Category{}).
		Select("slug").
		Where("slug = ? OR slug LIKE ?", baseSlug, baseSlug+"-%").
		Order("LENGTH(slug) DESC").
		Order("slug DESC").
		Limit(1).
		Pluck("slug", &slug).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil
	}

	return slug, nil
}
