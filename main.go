package main

import (
	"inventory-indra/db"
	"inventory-indra/handler"
	"inventory-indra/middleware"
	"inventory-indra/repositories"
	"time"

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:		[]string{"http://localhost:3000"},
		AllowMethods: 		[]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: 		[]string{"Origin", "Content-Type", "Authorization", "Accept", "x-user-id"},
		ExposeHeaders:		[]string{"Content-Length"},
		AllowCredentials: 	true,
    	MaxAge:           	12 * time.Hour,
	}))
	
	router.POST("/user", userHandler.Register)
	router.POST("/login", userHandler.Login)
	router.PATCH("/user", userHandler.PatchUser)
	router.GET("/user", middleware.TokenMiddleware, userHandler.GetUser)

	router.POST("/product", middleware.HandlerMiddleware, productHandler.CreateProduct)
	router.GET("/product", middleware.HandlerMiddleware, productHandler.GetProducts)
	router.DELETE("/product", middleware.HandlerMiddleware, productHandler.DeleteProduct)

	router.Run(":8000")
}
