package handler

import (
	"errors"
	"inventory-indra/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return  &UserHandler{repo: repo}
}

func (h *UserHandler) FindOneUser(ctx *gin.Context) {
	user, err := h.repo.FindUser()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			ctx.JSON(http.StatusNotFound, gin.H{
				"status": "failed",
				"message": "User is not found!",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error!",
			"status": "failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": user,
	})
}