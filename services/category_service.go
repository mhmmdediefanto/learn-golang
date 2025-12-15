package services

import (
	"fmt"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/repositories"
	"regexp"
	"strconv"

	"github.com/gosimple/slug"
)

type categoryService struct {
	repo repositories.CategoryRepository
}

type CategoryService interface {
	GetAll() ([]models.Category, error)
	GenerateUniqueSlug(title string) (string, error)
	Create(category *models.Category) (*models.Category, error)
	Delete(id uint) error
	Update(id uint, data *models.Category) (*models.Category, error)
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *categoryService) Create(category *models.Category) (*models.Category, error) {
	slug, err := s.GenerateUniqueSlug(category.Title)
	if err != nil {
		return nil, err
	}

	category.Slug = slug
	return s.repo.Create(category)
}

func (s *categoryService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *categoryService) Update(id uint, data *models.Category) (*models.Category, error) {

	if data.Slug == "" {
		slug, err := s.GenerateUniqueSlug(data.Title)
		if err != nil {
			return nil, err
		}
		data.Slug = slug
	}

	return s.repo.Update(data, id)
}

func (s *categoryService) GenerateUniqueSlug(title string) (string, error) {
	baseSlug := slug.Make(title)
	latestSlug, err := s.repo.FindLatestSlug(baseSlug)
	if err != nil || latestSlug == "" {
		return baseSlug, err
	}

	// kalau ada slug yang mirip, tambahkan angka di belakangnya
	re := regexp.MustCompile(`-(\d+)$`)

	// cek apakah latestSlug punya angka di belakang
	match := re.FindStringSubmatch(latestSlug)

	// jika ada, increment angkanya
	if len(match) == 2 {
		n, _ := strconv.Atoi(match[1])
		return fmt.Sprintf("%s-%d", baseSlug, n+1), nil
	}

	return baseSlug + "-2", nil

}
