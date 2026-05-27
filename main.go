package main

import (
	"inventory-indra/db"
	"inventory-indra/handler"
	"inventory-indra/middleware"
	"inventory-indra/repositories"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Server is starting...")
	db := db.Connect()

	userRepo := repositories.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	productRepo := repositories.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepo)

	stockRepo := repositories.NewStockRepository(db)
	stockHandler := handler.NewStockHandler(stockRepo)

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionHandler := handler.NewTransactionHandler(transactionRepo)

	statsRepo := repositories.NewStatsRepository(db)
	statsHandler := handler.NewStatsHandler(statsRepo)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://inventory-indra.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "x-user-id"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/user", userHandler.Register)
	router.POST("/login", userHandler.Login)
	router.PATCH("/user", userHandler.PatchUser)
	router.GET("/user", middleware.TokenMiddleware, userHandler.GetUser)

	router.POST("/product", middleware.HandlerMiddleware, productHandler.CreateProduct)
	router.GET("/product", middleware.HandlerMiddleware, productHandler.GetProducts)
	router.DELETE("/product", middleware.HandlerMiddleware, productHandler.DeleteProduct)
	router.PATCH("/product", middleware.HandlerMiddleware, productHandler.PatchProduct)

	router.GET("/stock", middleware.HandlerMiddleware, stockHandler.GetDataStocks)

	router.POST("/transaction", middleware.HandlerMiddleware, transactionHandler.CreateTransaction)
	router.GET("/transaction", middleware.HandlerMiddleware, transactionHandler.GetTransaction)

	router.GET("/stats", middleware.HandlerMiddleware, statsHandler.GetStatistik)
	router.GET("/stats/medicine", middleware.HandlerMiddleware, statsHandler.GetDataProductExcel)
	router.GET("/stats/reports", middleware.HandlerMiddleware, statsHandler.GetDataReportsExcel)
	router.Run(":8080")
}
