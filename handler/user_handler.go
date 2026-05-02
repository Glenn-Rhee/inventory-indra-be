package handler

import (
	"inventory-indra/model"
	"inventory-indra/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return  &UserHandler{repo: repo}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var reqBody model.CreateUser

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Bad request! Fill username and password!",
		})
		return
	}

	data, err := h.repo.LoginUser(reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "successfully login",
		"data": data,
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

func (h *UserHandler) GetUser(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Successfully get one user!",
	})
}

func (h *UserHandler) PatchUser(ctx *gin.Context){
	var reqBody model.UpdateUser

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Bad request! Fill id user and image url!",
		})
		return
	}

	if reqBody.Id == "" || reqBody.ImageUrl == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Bad request! Fill id user and image url!",
		})
		return
	}

	err, code := h.repo.PatchUser(reqBody)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status": "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Successfully update user",
	})	
}