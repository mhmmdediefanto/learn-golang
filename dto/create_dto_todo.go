package dto

type CreateDtoTodo struct {
	Title  string `json:"title" binding:"required,min=3"`
	Status bool   `json:"status"`
}
