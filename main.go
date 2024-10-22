package main

import (
	"context"
	"fmt"
	"my-budgety-pub-sub/services"
	"my-budgety-pub-sub/utils"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
)

func GetMessage() {
	envVariables := utils.GetEnvVariables()

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, envVariables.ProjectName)
	if err != nil {
		fmt.Println("Error creating pubsub client")
		panic(err)
	}
	defer client.Close()

	sub := client.Subscription(envVariables.SubscriptionId)

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})
	if err != nil {
		fmt.Println("Error receiving message")
		fmt.Print(err)
	}
	fmt.Println("Received", received, "messages")
}

func main() {
	projectName, err := services.GetSecretValue("project_name")
	if err != nil {
		panic("Error getting project name")
	}

	subscriptionName, err := services.GetSecretValue("pub-sub-subscription-id")
	if err != nil {
		panic("Error getting subscription name")
	}

	fmt.Printf("Project Name: %s\n", projectName)
	fmt.Printf("Subscription Name: %s\n", subscriptionName)
}
