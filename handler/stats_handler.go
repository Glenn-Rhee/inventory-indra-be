package handler

import (
	"inventory-indra/helper"
	"inventory-indra/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	repo *repositories.StatsRepository
}

func NewStatsHandler(repo *repositories.StatsRepository) *StatsHandler {
	return &StatsHandler{repo: repo}
}

func (h *StatsHandler) GetStatistik(ctx *gin.Context) {
	rangeTypeQuery, isExist := ctx.GetQuery("rangeType")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Please fill range type",
		})
		return
	}

	rangeType, err := strconv.Atoi(rangeTypeQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Please fill range type properly!",
		})
		return
	}

	if rangeType != 90 && rangeType != 30 && rangeType != 7 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Please fill range type only between 90, 30, or 7",
		})
		return
	}

	data, err, code := h.repo.GetStatistik(rangeType)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status":  "success",
		"message": "Successfully get data statistik",
		"data":    data,
	})
}

func (h *StatsHandler) GetDataProductExcel(ctx *gin.Context) {
	data, err, code := h.repo.GetDataProduct()
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	err = helper.TransactionsConvertToExcel(data, ctx)

	if err != nil {
		log.Println("Error create transactions file:", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "An error while create excel file",
		})

		return
	}
}

func (h *StatsHandler) GetDataReportsExcel(ctx *gin.Context) {
	data, err, code := h.repo.GetDataReports()
	if err != nil {
		ctx.JSON(code, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	err = helper.ReportsConvertToExcel(data, ctx)
	if err != nil {
		log.Println("Error create reports file:", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "An error while create excel file",
		})
	}
}
