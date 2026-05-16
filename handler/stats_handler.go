package handler

import (
	"inventory-indra/repositories"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	repo *repositories.StatsRepository
}

func NewStatsHandler(repo *repositories.StatsRepository) *StatsHandler {
	return &StatsHandler{repo: repo}
}

func (h *StatsHandler) GetStatistik(ctx *gin.Context) {
	data, err, code := h.repo.GetStatistik()
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