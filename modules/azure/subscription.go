package azure

import (
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-06-01/subscriptions"
)

// GetSubscriptionClient is a helper function that will setup an Azure Subscription client on your behalf
func GetSubscriptionClient() (*subscriptions.Client, error) {
	// Create a Subscription client
	client, err := CreateSubscriptionsClientE()
	if err != nil {
		return nil, err
	}

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	client.Authorizer = *authorizer
	return &client, nil
}
