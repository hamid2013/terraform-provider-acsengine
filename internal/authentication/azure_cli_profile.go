package authentication

import (
	"strings"

	"fmt"

	"github.com/Azure/go-autorest/autorest/azure/cli"
)

// AzureCLIProfile returns a go-autorest CLI profile struct
type AzureCLIProfile struct {
	cli.Profile
}

// FindDefaultSubscriptionID finds default subscription ID out of IDs
func (a AzureCLIProfile) FindDefaultSubscriptionID() (string, error) {
	for _, subscription := range a.Subscriptions {
		if subscription.IsDefault {
			return subscription.ID, nil
		}
	}

	return "", fmt.Errorf("no subscription was marked as default in the Azure Profile")
}

// FindSubscription finds subscription specified by given ID
func (a AzureCLIProfile) FindSubscription(subscriptionID string) (*cli.Subscription, error) {
	for _, subscription := range a.Subscriptions {
		if strings.EqualFold(subscription.ID, subscriptionID) {
			return &subscription, nil
		}
	}

	return nil, fmt.Errorf("subscription %q was not found in your Azure CLI credentials. Please verify it exists in `az account list`", subscriptionID)
}
