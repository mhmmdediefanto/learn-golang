package dto

type UpdateCategoryDto struct {
	Title string `json:"title" binding:"required"`
}
