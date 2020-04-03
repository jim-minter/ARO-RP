package billing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	azstorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/env"
)

const (
	tenantIDMSFT string = "72f988bf-86f1-41af-91ab-2d7cd011db47"
	tenantIDAME  string = "33e01921-4d64-4f8c-a055-5bdaffd5e33d"
)

// isSubscriptionRegisteredToE2E returns true if the subscription is having the
// "Microsoft.ContainerService/SaveAROTestConfig" feature registered
func IsSubscriptionRegisteredToE2E(sub *api.SubscriptionProperties) bool {
	if sub.TenantID == tenantIDMSFT || sub.TenantID == tenantIDAME {
		for _, feature := range sub.RegisteredFeatures {
			if feature.Name == api.SubscriptionFeatureE2E && feature.State == "Registered" {
				return true
			}
		}
	}
	return false
}

func CreateOrUpdateE2Eblob(ctx context.Context, env env.Interface, doc api.BillingDocument, resource azure.Resource) error {
	key, err := env.E2EStorageAccountKey(ctx)
	if err != nil {
		return err
	}
	client, err := azstorage.NewBasicClient(env.E2EStorageAccountName(), key)
	if err != nil {
		return err
	}
	blobclient := client.GetBlobService()

	containerName := strings.ToLower("bill-" + doc.Billing.Location + "-" + resource.ResourceGroup + "-" + resource.ResourceName)
	if len(containerName) > 63 {
		containerName = containerName[:62] // maximum characters allowed is 63
	}

	containerRef := blobclient.GetContainerReference(containerName)
	_, err = containerRef.CreateIfNotExists(nil)
	if err != nil {
		return fmt.Errorf("Error creating container : %s", err.Error())
	}

	blobRef := containerRef.GetBlobReference("billingentity")
	b, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	return blobRef.CreateBlockBlobFromReader(bytes.NewReader(b), nil)
}
