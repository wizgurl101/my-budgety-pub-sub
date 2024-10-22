package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	GCPAppCred     string
	ProjectName    string
	ProjectId      string
	SubscriptionId string
}

func GetEnvVariables() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return EnvVariables{
		GCPAppCred:     os.Getenv("SERVICE_ACCOUNT_PATH"),
		ProjectName:    os.Getenv("PROJECT_NAME"),
		ProjectId:      os.Getenv("PROJECT_ID"),
		SubscriptionId: os.Getenv("SUBSCRIPTION_ID"),
	}
}
