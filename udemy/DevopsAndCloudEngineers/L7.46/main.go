package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/wardviaene/golang-for-devops-course/ssh-demo"
)

const (
	location                 = "eastus"
	resourceGroupName        = "go-demo"
	virtualNetworkName       = "go-demo-vnet"
	subnetName               = "go-demo-subnet"
	publicIPAdressesName     = "go-demo-pip"
	networkSecurityGroupName = "go-demo-nsg"
	networkInterfaceFaceName = "go-demo-nic"
	iPConfigurationName      = "go-demo-ipconfig"
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
	// Create resource group
	if err := createResourceGroup(ctx, cred, subscriptionID, pubKey); err != nil {
		return err
	}

	// Create virtual network if not exists
	vnet, err := findVnet(ctx, cred, subscriptionID)
	if err != nil {
		return err
	}

	if vnet == nil {
		if _, err := createVirtualNetwork(ctx, cred, subscriptionID); err != nil {
			return err
		}
	}

	// Create subnet
	subnet, err := createSubnet(ctx, cred, subscriptionID)
	if err != nil {
		return err
	}

	// Create public IP address
	publicIPAddress, err := createPublicIpClient(ctx, cred, subscriptionID)
	if err != nil {
		return err
	}

	// Create network security group
	networkSecurityGroup, err := createNetworkSecurityGroup(ctx, cred, subscriptionID)
	if err != nil {
		return err
	}

	// Create network interface client
	if _, err := createNetworkInterFaceClient(ctx, cred, subscriptionID, networkSecurityGroup, subnet, publicIPAddress); err != nil {
		return err
	}

	return nil
}

func createResourceGroup(ctx context.Context, cred azcore.TokenCredential, subscriptionID, pubKey string) error {
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

func createVirtualNetwork(ctx context.Context, cred azcore.TokenCredential, subscriptionID string) (*armnetwork.VirtualNetwork, error) {
	virtualNetworkClient, err := armnetwork.NewVirtualNetworksClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}

	pollerResp, err := virtualNetworkClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		virtualNetworkName,
		armnetwork.VirtualNetwork{
			Location: to.Ptr(location),
			Properties: &armnetwork.VirtualNetworkPropertiesFormat{
				AddressSpace: &armnetwork.AddressSpace{
					AddressPrefixes: []*string{
						to.Ptr("10.1.0.0/16"),
					},
				},
			},
		},
		nil)

	if err != nil {
		return nil, err
	}

	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &resp.VirtualNetwork, nil
}

func createSubnet(ctx context.Context, cred azcore.TokenCredential, subscriptionID string) (*armnetwork.Subnet, error) {
	subnetsClient, err := armnetwork.NewSubnetsClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}

	pollerResp, err := subnetsClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		virtualNetworkName,
		subnetName,
		armnetwork.Subnet{
			Properties: &armnetwork.SubnetPropertiesFormat{
				AddressPrefix: to.Ptr("10.1.0.0/24"),
			},
		},
		nil,
	)

	if err != nil {
		return nil, err
	}

	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &resp.Subnet, nil
}

func createPublicIpClient(ctx context.Context, cred azcore.TokenCredential, subscriptionID string) (*armnetwork.PublicIPAddress, error) {
	publicIPAdressesClient, err := armnetwork.NewPublicIPAddressesClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}

	pollerResp, err := publicIPAdressesClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		publicIPAdressesName,
		armnetwork.PublicIPAddress{
			Location: to.Ptr(location),
			Properties: &armnetwork.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
			},
		},
		nil,
	)

	if err != nil {
		return nil, err
	}

	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &resp.PublicIPAddress, nil
}

func createNetworkSecurityGroup(ctx context.Context, cred azcore.TokenCredential, subscriptionID string) (*armnetwork.SecurityGroup, error) {
	securityGroupsClient, err := armnetwork.NewSecurityGroupsClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}

	pollerResp, err := securityGroupsClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		networkSecurityGroupName,
		armnetwork.SecurityGroup{
			Location: to.Ptr(location),
			Properties: &armnetwork.SecurityGroupPropertiesFormat{
				SecurityRules: []*armnetwork.SecurityRule{
					{
						Name: to.Ptr("allow-ssh"),
						Properties: &armnetwork.SecurityRulePropertiesFormat{
							SourceAddressPrefix:      to.Ptr("0.0.0.0/0"),
							SourcePortRange:          to.Ptr("*"),
							DestinationAddressPrefix: to.Ptr("0.0.0.0/0"),
							DestinationPortRange:     to.Ptr("22"),
							Protocol:                 to.Ptr(armnetwork.SecurityRuleProtocolTCP),
							Access:                   to.Ptr(armnetwork.SecurityRuleAccessAllow),
							Description:              to.Ptr("Allow SSH traffic on all ports"),
							Direction:                to.Ptr(armnetwork.SecurityRuleDirectionInbound),
							Priority:                 to.Ptr(int32(1001)),
						},
					},
				},
			},
		},
		nil,
	)

	if err != nil {
		return nil, err
	}

	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &resp.SecurityGroup, nil
}

func createNetworkInterFaceClient(
	ctx context.Context,
	cred azcore.TokenCredential,
	subscriptionID string,
	networkSecurityGroup *armnetwork.SecurityGroup,
	subnet *armnetwork.Subnet,
	publicIPAdress *armnetwork.PublicIPAddress,
) (*armnetwork.Interface, error) {
	interfacesClient, err := armnetwork.NewInterfacesClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}

	pollerResp, err := interfacesClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		networkInterfaceFaceName,
		armnetwork.Interface{
			Location: to.Ptr(location),
			Properties: &armnetwork.InterfacePropertiesFormat{
				NetworkSecurityGroup: &armnetwork.SecurityGroup{
					ID: networkSecurityGroup.ID,
				},
				IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
					{
						Name: to.Ptr(iPConfigurationName),
						Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
							Subnet: &armnetwork.Subnet{
								ID: subnet.ID,
							},
							PublicIPAddress: &armnetwork.PublicIPAddress{
								ID: publicIPAdress.ID,
							},
						},
					},
				},
			},
		},
		nil,
	)

	if err != nil {
		return nil, err
	}

	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &resp.Interface, nil
}

func findVnet(ctx context.Context, cred azcore.TokenCredential, subscriptionID string) (*armnetwork.VirtualNetwork, error) {
	virtualNetworksClient, err := armnetwork.NewVirtualNetworksClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}
	// TODO: Find a better way to check if the vnet exists
	resp, err := virtualNetworksClient.Get(ctx, resourceGroupName, virtualNetworkName, nil)
	if err != nil {
		var errResponse *azcore.ResponseError
		if errors.As(err, &errResponse) && errResponse.ErrorCode == "ResourceNotFound" {
			return nil, nil
		}
		return nil, err
	}
	return &resp.VirtualNetwork, nil
}

type MyPollerResp interface {
	*runtime.Poller[armnetwork.VirtualNetworksClientCreateOrUpdateResponse] |
		*runtime.Poller[armnetwork.SubnetsClientCreateOrUpdateResponse] |
		*runtime.Poller[armnetwork.PublicIPAddressesClientCreateOrUpdateResponse] |
		*runtime.Poller[armnetwork.SecurityGroupsClientCreateOrUpdateResponse] |
		*runtime.Poller[armnetwork.InterfacesClientCreateOrUpdateResponse]

		PollUntilDone(ctx context.Context, options *runtime.PollUntilDoneOptions) (interface{}, error)
}

func MyFunc[T MyPollerResp](ctx context.Context, pollerResp T) (interface{} ,error){
	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
