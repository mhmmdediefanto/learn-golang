package services

import (
	"fmt"
	"go-bakcend-todo-list/dto"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/pkg/apperror"
	"go-bakcend-todo-list/repositories"
	"regexp"
	"strconv"

	"github.com/gosimple/slug"
)

type ArticleService interface {
	GetAll() ([]models.Article, error)
	Create(article *models.Article) (*models.Article, error)
	Delete(id uint, userID uint) error
	Update(userID uint, id uint, article dto.UpdateArticleDto) (*models.Article, error)
}

type articleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{
		repo: repo,
	}
}

func (s *articleService) GetAll() ([]models.Article, error) {
	return s.repo.GetAll()
}

func (s *articleService) Create(data *models.Article) (*models.Article, error) {
	slug, err := s.GenerateUniqueSlug(data.Title)
	if err != nil {
		return nil, err
	}

	data.Slug = slug

	return s.repo.Create(data)
}

func (s *articleService) GenerateUniqueSlug(title string) (string, error) {
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

func (s *articleService) Delete(id uint, userID uint) error {
	return s.repo.DeleteByUser(id, userID)
}

func (s *articleService) Update(userID uint, id uint, article dto.UpdateArticleDto) (*models.Article, error) {
	articleResult, err := s.repo.FindById(id)
	if err != nil {
		return nil, apperror.NotFound("Data Tidak di temukan")
	}

	if userID != articleResult.UserID {
		return nil, apperror.Forbidden("Tidak punya akses ke artikel ini")
	}

	if articleResult.Title != article.Title {

		baseSlug, err := s.GenerateUniqueSlug(article.Title)
		if err != nil {
			return nil, err
		}

		articleResult.Slug = baseSlug
	}

	articleResult.Title = article.Title
	articleResult.Content = article.Content

	if article.Status != "" {
		articleResult.Status = article.Status
	}

	return s.repo.UpdateByUser(id, userID, articleResult)
}
