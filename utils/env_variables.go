package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	ProjectName    string
	GCPAppCred     string
	SubscriptionId string
}

func GetEnvVariables() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return EnvVariables{
		ProjectName:    os.Getenv("PROJECT_NAME"),
		GCPAppCred:     os.Getenv("SERVICE_ACCOUNT_PATH"),
		SubscriptionId: os.Getenv("SUBSCRIPTION_ID"),
	}
}
