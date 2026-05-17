package handler

import (
	"inventory-indra/repositories"
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
			"status": "failed",
			"message": "Please fill range type",
		})
		return
	}

	rangeType, err := strconv.Atoi(rangeTypeQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Please fill range type properly!",
		})
		return
	}

	if rangeType != 90 && rangeType != 30 && rangeType != 7 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Please fill range type only between 90, 30, or 7",
		})
		return
	}

	data, err, code := h.repo.GetStatistik(rangeType)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status": "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status": "success",
		"message": "Successfully get data statistik",
		"data": data,
	})
}