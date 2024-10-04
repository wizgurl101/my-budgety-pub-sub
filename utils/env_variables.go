package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	TestVariable string
}

func GetEnvVariables() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return EnvVariables{
		TestVariable: os.Getenv("TEST_VARIABLE"),
	}
}
