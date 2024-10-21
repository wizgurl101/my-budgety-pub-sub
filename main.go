package main

import (
	"context"
	"fmt"
	"my-budgety-pub-sub/utils"

	"cloud.google.com/go/pubsub"
)

func main() {
	envVariables := utils.GetEnvVariables()
	fmt.Println("GCP App Credential:", envVariables.GCPAppCred)

	context := context.Background()
	client, err := pubsub.NewClient(context, envVariables.ProjectName)
	if err != nil {
		fmt.Println("Error creating pubsub client")
		panic(err)
	}
	defer client.Close()

}
