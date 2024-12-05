package services

import (
	"context"
	"fmt"

	billing "cloud.google.com/go/billing/apiv1"
	billingpb "cloud.google.com/go/billing/apiv1/billingpb"
)

func disableBilling(projectID string, billingAccountName string) (*billingpb.ProjectBillingInfo, error) {
	ctx := context.Background()

	client, err := billing.NewCloudBillingClient(ctx)
	if err != nil {
		fmt.Printf("billing.NewCloudBillingClient: %v", err)
		return nil, err
	}
	defer client.Close()

	// setting the Billing Account Name to empty string will disable billing
	billingInfo := &billingpb.ProjectBillingInfo{
		Name:               "projects/" + projectID + "/billingInfo",
		ProjectId:          projectID,
		BillingAccountName: "",
		BillingEnabled:     false,
	}

	request := &billingpb.UpdateProjectBillingInfoRequest{
		Name:               billingAccountName,
		ProjectBillingInfo: billingInfo,
	}

	response, err := client.UpdateProjectBillingInfo(ctx, request)
	if err != nil {
		fmt.Printf("client.UpdateProjectBillingInfo: %v", err)
		return nil, err
	}

	fmt.Printf("Billing disabled for project: %s\n", projectID)
	return response, nil
}
