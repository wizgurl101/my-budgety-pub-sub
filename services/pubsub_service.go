package services

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func DiableBillingIfBudgetIsReach(projectName string, subscriptionId string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectName)
	if err != nil {
		panic(
			fmt.Sprintf("Error creating pubsub client: %v", err),
		)
	}
	defer client.Close()

	sub := client.Subscription(subscriptionId)

	// var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		type MessageData struct {
			CostAmount   float32 `json:"costAmount"`
			BudgetAmount float32 `json:"budgetAmount"`
		}

		var data MessageData
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			fmt.Printf("Error unmarshalling message data: %v\n", err)
			msg.Nack()
			return
		}

		var spendAmount float32 = data.CostAmount
		var budgetAmount float32 = data.BudgetAmount

		if spendAmount == budgetAmount {
			fmt.Printf("disable billing")

			// send email notifcations

		} else {
			fmt.Printf("spend amount is less than budget amount")
		}

		// atomic.AddInt32(&received, 1)
		// msg.Ack()
	})
	if err != nil {
		panic(
			fmt.Sprintf("Error receiving message: %v", err),
		)
	}
}
