package services

import (
	"context"
	"fmt"
	"my-budgety-pub-sub/utils"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func GetSecretValue(name string) (string, error) {
	envVariables := utils.GetEnvVariables()
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		fmt.Printf("failed to create secret manager client: %v\n", err)
		return "", err
	}
	defer client.Close()

	request := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", envVariables.ProjectId, name),
	}

	result, err := client.AccessSecretVersion(ctx, request)
	if err != nil {
		fmt.Printf("failed to access secret version: %v\n", err)
		return "", err
	}

	return string(result.Payload.Data), nil
}
