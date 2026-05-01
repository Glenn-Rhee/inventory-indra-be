package handler

import (
	"errors"
	"inventory-indra/model"
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

func (h *UserHandler) Register(ctx *gin.Context) {
	var reqBody model.CreateUser

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Bad request! Fill username and password!",
		})
		return
	}

	if reqBody.Username == "" || reqBody.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Bad request! Fill username and password!",
		})
		return
	}

	err := h.repo.CreateUser(reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Successfully signup!",
	})
}