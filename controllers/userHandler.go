package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/SMS-Verification/db"
	"github.com/waltertaya/SMS-Verification/models"
	"github.com/waltertaya/SMS-Verification/utils"
)

func LoginUser(ctx *gin.Context) {
	var userBody struct {
		Username string
		Password string
	}

	ctx.Bind(&userBody)

	user := models.User{}

	result := db.DB.Where("username = ?", userBody.Username).First(&user)

	if result.Error != nil {
		// utils.LogErrors(result.Error) -> causes panic (exiting) : bug
		// fmt.Println("user not found")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// passCheck := utils.VerifyPassword(user.Password, user.Password) -> comparing hash to itself : bug
	passCheck := utils.VerifyPassword(userBody.Password, user.Password)

	if !passCheck {
		// fmt.Println("Password check", passCheck)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	if !user.Verified {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Verify the user using phone number",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User login successfully",
		"user": user,
	})
}

func RegisterUser(ctx *gin.Context) {
	var userBody struct {
		Username string
		Name     string
		Email    string
		Password string
	}

	ctx.Bind(&userBody)

	hashedPassword, err := utils.HashPassword(userBody.Password)

	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"error": "Error registering the user",
		})
		return
	}

	id := utils.GenId()

	user := models.User{
		Username: userBody.Username,
		Name:     userBody.Name,
		Email:    userBody.Email,
		Password: hashedPassword,
		Role:     "normal", // normal, admin
		Verified: false,
		ID:       id,
	}

	user_exist := db.DB.Where("username = ? OR email = ?", userBody.Username, userBody.Email).First(&user)

	if user_exist.RowsAffected > 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "User already exist",
		})
		return
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "Error registering the user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "User registered successfully",
		"user": user,
	})
}
