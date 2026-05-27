package handler

import (
	"inventory-indra/model"
	"inventory-indra/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	repo *repositories.TransactionRepository
}

func NewTransactionHandler(repo *repositories.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repo: repo}
}

func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var reqBody model.CreateTransaction

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Bad request! Fill field of product properly!",
		})
		return
	}

	valUserId, isExist := ctx.Get("userId")
	if !isExist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Unauthorized! Please make sure you've been login!",
		})
		return
	}

	userId, ok := valUserId.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Invalid user id!",
		})
		return
	}

	err, code := h.repo.CreateTransaction(reqBody, userId)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status":  "success",
		"message": "Successfully create transaction!",
	})
}

func (h *TransactionHandler) GetTransaction(ctx *gin.Context) {
	limitQuery, isExist := ctx.GetQuery("limit")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Please fill limit query!",
		})
		return
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Please fill limit properly!",
		})
		return
	}

	pageQuery, isExist := ctx.GetQuery("page")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Please fill page query!",
		})
		return
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Please fill page table properly!",
		})
		return
	}

	filter, _ := ctx.GetQuery("filter")

	data, err, code := h.repo.GetTransaction(repositories.GetTransactionParams{
		Limit:  limit,
		Page:   page,
		Filter: filter,
	})

	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status":  "success",
		"message": "Successfully get data transaction",
		"data":    data,
	})
}
