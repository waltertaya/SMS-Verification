package initializers

import (
	"github.com/joho/godotenv"
	"github.com/waltertaya/SMS-Verification/utils"
)

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		utils.LogErrors(err)
	}
}
