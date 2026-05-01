package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenMiddleware(ctx *gin.Context ) {
	token := ctx.GetHeader("Authorization")
	log.Println("Token Auth:", token)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"message": "Please fill header Authorization!",
		})
		ctx.Abort()
		return
	}
	parts := strings.Split(token, " ")
	if len(parts) < 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"message": "Invalid Token Format!",
		})
		ctx.Abort()
		return
	}

	token = parts[1]
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"message": "Unathorized! Please login to your account first!",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}