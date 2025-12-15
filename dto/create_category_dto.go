package dto

type CreateCategoryDto struct {
	Title string `json:"title" binding:"required"`
}
