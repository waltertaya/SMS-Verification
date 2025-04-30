package db

import (
	"fmt"
	"os"

	"github.com/waltertaya/SMS-Verification/initializers"
	"github.com/waltertaya/SMS-Verification/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	initializers.LoadEnv()

	var err error

	dsn := os.Getenv("DB_URI")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err == nil {
		fmt.Println("Mysql Database connected successfully")
	} else {
		utils.LogErrors(err)
	}
}
