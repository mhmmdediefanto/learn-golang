package controllers

import (
	"go-bakcend-todo-list/dto"
	"go-bakcend-todo-list/enums"
	"go-bakcend-todo-list/models"
	"go-bakcend-todo-list/pkg/apperror"
	"go-bakcend-todo-list/services"
	"go-bakcend-todo-list/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	service services.ArticleService
}

func NewArticleController(service services.ArticleService) *ArticleController {
	return &ArticleController{
		service: service,
	}
}

func (c *ArticleController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		utils.Error(ctx, 500, "Gagal Mengambil Article", err)
		return
	}

	utils.Success(ctx, "Retrieved Article Successfully", data)
}

func (c *ArticleController) Create(ctx *gin.Context) {
	UserID := utils.MustCurrentUser(ctx)

	var req dto.CreateArticleDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	if req.Status == "" {
		req.Status = string(enums.StatusDraft)
	}

	article := &models.Article{
		Title:   req.Title,
		Content: req.Content,
		Status:  req.Status,
		UserID:  UserID,
	}

	data, err := c.service.Create(article)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "gagal membuat articel", err)
		return
	}

	utils.Created(ctx, "Endpoint create article", data)
}

func (c *ArticleController) Delete(ctx *gin.Context) {
	userID := utils.MustCurrentUser(ctx)
	idParams, err := utils.ParamUint(ctx, "id")
	if err != nil {
		utils.Error(ctx, 400, "ID tidak valid", err)
		return

	}
	if err := c.service.Delete(idParams, userID); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal Menghapus Article", err)
		return
	}
	utils.Success(ctx, "Article Berhasil di hapus", nil)

}

func (c *ArticleController) Update(ctx *gin.Context) {
	// Ambil user ID dari context (biasanya dari JWT / middleware auth)
	userID := utils.MustCurrentUser(ctx)

	// Ambil param ID dan validasi
	idParams, err := utils.ParamUint(ctx, "id")
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "ID tidak valid", err)
		return
	}

	// Bind request body
	var reqUpdateArticle dto.UpdateArticleDto
	if err := ctx.ShouldBindJSON(&reqUpdateArticle); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// Panggil service
	result, err := c.service.Update(userID, idParams, reqUpdateArticle)
	if err != nil {
		// Kalau error dari service adalah AppError
		if appErr, ok := err.(*apperror.AppError); ok {
			utils.Error(ctx, appErr.Code, appErr.Message, appErr.Err)
			return
		}

		// Fallback (error tidak terduga)
		utils.Error(ctx, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	// Success response
	utils.Success(ctx, "Update article berhasil", result)
}
