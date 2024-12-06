package main

import (
	"fmt"
	"my-budgety-pub-sub/services"
)

func main() {
	projectNameChan := make(chan string)
	subscriptionNameChan := make(chan string)
	billingAccountNameChan := make(chan string)
	errChan := make(chan error)

	go func() {
		projectName, err := services.GetSecretValue("project_name")
		if err != nil {
			errChan <- fmt.Errorf("error getting project name: %v", err)
			return
		}
		projectNameChan <- projectName
	}()

	go func() {
		subscriptionName, err := services.GetSecretValue("pub-sub-subscription-id")
		if err != nil {
			errChan <- fmt.Errorf("error getting subscription name: %v", err)
			return
		}
		subscriptionNameChan <- subscriptionName
	}()

	go func() {
		billingAccountName, err := services.GetSecretValue("billing_account_name")
		if err != nil {
			errChan <- fmt.Errorf("error getting billing account name: %v", err)
			return
		}
		billingAccountNameChan <- billingAccountName
	}()

	var projectName, subscriptionName, billingAccountName string
	for i := 0; i < 2; i++ {
		select {
		case err := <-errChan:
			panic(err)
		case projectName = <-projectNameChan:
		case subscriptionName = <-subscriptionNameChan:
		case billingAccountName = <-billingAccountNameChan:
		}
	}

	disableBilling, err = services.isSpendingGreaterThanBudget(projectName, subscriptionName, billingAccountName)
}
