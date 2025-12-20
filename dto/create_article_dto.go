package dto

type CreateArticleDto struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status" binding:"omitempty,oneof=draft published archived"`
}
