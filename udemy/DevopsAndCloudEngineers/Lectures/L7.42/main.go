package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/wardviaene/golang-for-devops-course/ssh-demo"
)

const (
	location          = "eastus"
	resourceGroupName = "go-demo"
)

func main() {
	var (
		token          azcore.TokenCredential
		pubKey         string
		err            error
		subscriptionID string
	)

	ctx := context.Background()
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		fmt.Println("AZURE SUBSCRIPTION ID not set")
		os.Exit(1)
	}

	if pubKey, err = generateKey(); err != nil {
		fmt.Printf("generateKey() error: %s\n", err)
		os.Exit(1)
	}
	if token, err = getToken(); err != nil {
		fmt.Printf("getToken() error: %s\n", err)
		os.Exit(1)
	}
	if err = launchInstance(ctx, token, subscriptionID, pubKey); err != nil {
		fmt.Printf("launchInstance() error: %s\n", err)
		os.Exit(1)
	}
}

func generateKey() (string, error) {
	var (
		privateKey []byte
		publicKey  []byte
		err        error
	)

	if privateKey, publicKey, err = ssh.GenerateKeys(); err != nil {
		return "", err
	}

	if err = os.WriteFile("myKey.pem", privateKey, 0600); err != nil {
		return "", err
	}

	if err = os.WriteFile("myKey.pub", publicKey, 0644); err != nil {
		return "", err
	}

	return string(publicKey), nil
}

func getToken() (azcore.TokenCredential, error) {
	token, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func launchInstance(ctx context.Context, cred azcore.TokenCredential, subscriptionID, pubKey string) error {
	
	if err := createResourceClient(ctx, cred, subscriptionID, pubKey); err != nil {
		return err
	}
	return nil
}

func createResourceClient(ctx context.Context, cred azcore.TokenCredential, subscriptionID, pubKey string) error {
	options := &arm.ClientOptions{}

	resourcesGroupClient, err := armresources.NewResourceGroupsClient(subscriptionID, cred, options)
	if err != nil {
		return err
	}

	resourceGroupParams := armresources.ResourceGroup{
		Location: to.Ptr(location),
	}

	resourceGroupResp, err := resourcesGroupClient.CreateOrUpdate(ctx, resourceGroupName, resourceGroupParams, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Resource group created: %s\n", *resourceGroupResp.ID)

	return nil
}
