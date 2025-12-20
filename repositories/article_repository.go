package repositories

import (
	"errors"
	"go-bakcend-todo-list/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetAll() ([]models.Article, error)
	FindLatestSlug(slug string) (string, error)
	Create(data *models.Article) (*models.Article, error)
	FindById(id uint) (*models.Article, error)
	DeleteByUser(id uint, userID uint) error
	UpdateByUser(id uint, userID uint, article *models.Article) (*models.Article, error)
}
type articleRepository struct {
	db *gorm.DB
}

// contructor
func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

// method get all
func (r *articleRepository) GetAll() ([]models.Article, error) {
	var articles []models.Article
	if err := r.db.Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

// method create
func (r *articleRepository) Create(data *models.Article) (*models.Article, error) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *articleRepository) FindLatestSlug(baseSlug string) (string, error) {
	var slug string

	err := r.db.Model(&models.Article{}).
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

func (r *articleRepository) FindById(id uint) (*models.Article, error) {
	var article models.Article

	err := r.db.First(&article, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("article tidak ditemukan")
		}
		return nil, err
	}

	return &article, nil
}

func (r *articleRepository) DeleteByUser(id uint, userID uint) error {
	result := r.db.
		Where("id = ? AND user_id = ?", id, userID).Delete(&models.Article{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("artikel tidak ditemukan atau bukan milik anda")
	}

	return nil
}

func (r *articleRepository) UpdateByUser(id uint, userID uint, article *models.Article) (*models.Article, error) {
	result := r.db.Model(&models.Article{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"title":   article.Title,
			"content": article.Content,
			"slug":    article.Slug,
			"status":  article.Status,
		})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("gagal update article")
	}

	return article, nil
}
