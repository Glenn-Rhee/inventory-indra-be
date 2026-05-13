package handler

import (
	"inventory-indra/model"
	"inventory-indra/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	repo *repositories.TransactionRepository
}

func NewTransactionHandler(repo *repositories.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repo: repo}
}

func (h *TransactionHandler) CreateTransaction(ctx *gin.Context){
	var reqBody model.CreateTransaction

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Bad request! Fill field of product properly!",
		})
		return
	}

	err, code := h.repo.CreateTransaction(reqBody)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status": "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status": "success",
		"message": "Successfully create transaction!",
	})
}