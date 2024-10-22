package main

import (
	"my-budgety-pub-sub/services"
)

func main() {
	projectName, err := services.GetSecretValue("project_name")
	if err != nil {
		panic("Error getting project name")
	}

	subscriptionName, err := services.GetSecretValue("pub-sub-subscription-id")
	if err != nil {
		panic("Error getting subscription name")
	}

	services.GetMessage(projectName, subscriptionName)
}
