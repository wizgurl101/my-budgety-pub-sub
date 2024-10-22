package services

import (
	"context"
	"fmt"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
)

func GetMessage(projectName string, subscriptionId string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectName)
	if err != nil {
		panic(
			fmt.Sprintf("Error creating pubsub client: %v", err),
		)
	}
	defer client.Close()

	sub := client.Subscription(subscriptionId)

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})
	if err != nil {
		panic(
			fmt.Sprintf("Error receiving message: %v", err),
		)
	}
	fmt.Println("Received", received, "messages")
}
