package handler

import (
	"inventory-indra/model"
	"inventory-indra/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo *repositories.ProductRepository
}

func NewProductHandler(repo *repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context){
	var reqBody model.CreateProduct

	if err := ctx.ShouldBindJSON(&reqBody); err !=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Bad request! Fill field of product properly!",
		})
		return
	}

	err, code := h.repo.CreateProduct(reqBody)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status": "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"message": "Successfully create product!",
	})
}