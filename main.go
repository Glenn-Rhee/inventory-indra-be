package main

import (
	"inventory-indra/db"
	"inventory-indra/handler"
	"inventory-indra/middleware"
	"inventory-indra/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.Connect()

	userRepo := repositories.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)
	productRepo := repositories.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepo)

	router := gin.Default()

	router.Use(cors.Default())
	
	router.POST("/user", userHandler.Register)
	router.POST("/login", userHandler.Login)
	router.PATCH("/user", userHandler.PatchUser)
	router.GET("/user", middleware.TokenMiddleware, userHandler.GetUser)

	router.POST("/product", productHandler.CreateProduct)
	router.Run(":8000")
}
