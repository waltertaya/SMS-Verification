package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/waltertaya/SMS-Verification/controllers"
	"github.com/waltertaya/SMS-Verification/db"
	"github.com/waltertaya/SMS-Verification/initializers"
)

func init() {
	initializers.LoadEnv()
	db.ConnectDB()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:80"}, // Allow frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "You have successfully reached the sms verification endpoints",
		})
	})

	r.POST("/api/v2/auth/login", controllers.LoginUser)
	r.POST("/api/v2/auth/register", controllers.RegisterUser)
	r.POST("/api/v2/auth/2fa/code", controllers.CodeRequest)
	r.POST("/api/v2/auth/2fa/verify", controllers.VerifyCode)

	r.Run()
}
