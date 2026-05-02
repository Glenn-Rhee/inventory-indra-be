package main

import (
	"inventory-indra/db"
	"inventory-indra/handler"
	"inventory-indra/middleware"
	"inventory-indra/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.Connect()
	userRepo := repositories.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)
	
	router := gin.Default()
	router.POST("/user", userHandler.Register)
	router.POST("/login", userHandler.Login)
	router.GET("/user", middleware.TokenMiddleware, userHandler.GetUser)
	router.Run(":8000")
}
