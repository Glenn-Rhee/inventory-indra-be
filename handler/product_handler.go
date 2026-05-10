package handler

import (
	"inventory-indra/model"
	"inventory-indra/repositories"
	"log"
	"net/http"
	"strconv"

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
		log.Println(ctx.Request.Body)
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

func (h *ProductHandler) GetProducts(ctx *gin.Context){
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


	dataProducts, err := h.repo.GetProducts(limit, page, filter)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"message": err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Successfully get data products!!",
		"data": dataProducts,
	})
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context){
	productId, isExist := ctx.GetQuery("productId")
	if !isExist || productId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Please fill product id!",
		})
		return
	}

	err, code := h.repo.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(code, gin.H{
			"status": "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"status": "success",
		"message": "Successfully Delete product!",
	})
}