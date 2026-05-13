package handler

import (
	"inventory-indra/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	repo *repositories.StockRepository
}

func NewStockHandler(repo *repositories.StockRepository) *StockHandler {
	return &StockHandler{repo: repo}
}

func (h *StockHandler) GetDataStocks(ctx *gin.Context){
	limitQuery, isExist := ctx.GetQuery("limit")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Please fill limit query!",
		})
		return
	}

 	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Please fill limit properly!",
		})
		return
	}

	pageQuery, isExist := ctx.GetQuery("page")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Please fill page query!",
		})
		return
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Please fill page table properly!",
		})
		return
	}

	filter, _ := ctx.GetQuery("filter")

	data, err, code := h.repo.GetStocks(repositories.GetStocksParams{Limit: limit, Page: page, Filter: filter})

	if err != nil {
		ctx.JSON(code, gin.H{
			"status": "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status": "success",
		"message": "Successfully get data stocks!",
		"data": data,
	})
}