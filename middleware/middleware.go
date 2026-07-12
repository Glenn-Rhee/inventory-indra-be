package middleware

import (
	"errors"
	"fmt"
	"inventory-indra/db"
	"inventory-indra/helper"
	"inventory-indra/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func HandlerMiddleware(ctx *gin.Context) {
	sessionUserId := ctx.GetHeader("x-user-id")
	if sessionUserId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Please fill header x-user-id",
		})
		ctx.Abort()
		return
	}

	var user model.User

	result := db.DB.Where("id = ?", sessionUserId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Unauthorized! Please make sure you've been login!",
		})
		ctx.Abort()
		return
	}

	ctx.Set("userId", user.Id)
	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	log.Println("Token Auth:", token)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Please fill header Authorization!",
		})
		ctx.Abort()
		return
	}
	parts := strings.Split(token, " ")
	if len(parts) < 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Invalid Token Format!",
		})
		ctx.Abort()
		return
	}

	token = parts[1]
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Unathorized! Please login to your account first!",
		})
		ctx.Abort()
		return
	}

	claims, err := helper.ValidateJWT(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenInvalidId) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Invalid Token!",
			})
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Token is expired!",
			})
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Your session is ended! Kindly login again.",
			})
		} else {
			fmt.Println("Error:", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Your token is invalid! Please login!",
			})
		}
		ctx.Abort()
		return
	}

	ctx.Set("userId", claims.Id)

	ctx.Next()
}

