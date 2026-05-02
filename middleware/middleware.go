package middleware

import (
	"errors"
	"fmt"
	"inventory-indra/helper"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	claims, err := helper.ValidateJWT(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenInvalidId) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"message": "Invalid Token!",
			})
		} else if errors.Is(err, jwt.ErrTokenExpired){
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"message": "Token is expired!",
			})
		} else if errors.Is(err, jwt.ErrTokenMalformed){
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"message": "Your session is ended! Kindly login again.",
			})
		} else {
			fmt.Println("Error:", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"message": "Your token is invalid! Please login!",
			})
		}
		ctx.Abort()
		return
	}

	ctx.Set("userId", claims.Id)

	ctx.Next()
}