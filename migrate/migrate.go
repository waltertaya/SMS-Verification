package main

import (
	"github.com/waltertaya/SMS-Verification/db"
	"github.com/waltertaya/SMS-Verification/initializers"
	"github.com/waltertaya/SMS-Verification/models"
)

func init() {
	initializers.LoadEnv()
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.User{})
}
