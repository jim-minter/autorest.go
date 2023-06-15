//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

var (
	err               error
	ctx               context.Context
	cred              azcore.TokenCredential
	letterRunes       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	armEndpoint       = "https://management.azure.com"
	fakeStepVar       = "signalrswaggertest4"
	resourceName      = "signalrswaggertest4"
	testPrefix        = generateAlphaNumericID("test", 6)
	location          = getEnv("LOCATION", "westus")
	resourceGroupName = getEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	subscriptionId    = getEnv("AZURE_SUBSCRIPTION_ID", "")
)

func main() {
	ctx = context.Background()
	cred, err = azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}
	createResourceGroup()
	prepare()
	scenario0Sample()
	scenario1Sample()
	deleteResourceGroup()
}

func prepare() {
	// From step Delete-proximity-placement-group
	proximityPlacementGroupsClient, err := armcompute.NewProximityPlacementGroupsClient(subscriptionId, cred, nil)
	if err != nil {
		panic(err)
	}
	_, err = proximityPlacementGroupsClient.Delete(ctx, resourceGroupName, resourceName, nil)
	if err != nil {
		panic(err)
	}
}

// Microsoft.SignalRService/Basic_CRUD
func scenario0Sample() {
	fakeScenarioVar := "signalrswaggertest5"
	resourceName := resourceName
	// From step Generate_Unique_Name
	template := map[string]any{
		"$schema":        "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
		"contentVersion": "1.0.0.0",
		"outputs": map[string]any{
			"name": map[string]any{
				"type":  "string",
				"value": "[variables('name').value]",
			},
			"resourceName": map[string]any{
				"type":  "string",
				"value": "[variables('name').value]",
			},
		},
		"resources": []any{},
		"variables": map[string]any{
			"name": map[string]any{
				"type": "string",
				"metadata": map[string]any{
					"description": "Name of the SignalR service.",
				},
				"value": "[concat('sw',uniqueString(resourceGroup().id))]",
			},
		},
	}
	deployment := armresources.Deployment{
		Properties: &armresources.DeploymentProperties{
			Template: template,
			Mode:     to.Ptr(armresources.DeploymentModeIncremental),
		},
	}
	deploymentExtend := createDeployment("Generate_Unique_Name", &deployment)
	name = deploymentExtend.Properties.Outputs.(map[string]interface{})["name"].(map[string]interface{})["value"].(string)
	resourceName = deploymentExtend.Properties.Outputs.(map[string]interface{})["resourceName"].(map[string]interface{})["value"].(string)

	// From step Create-or-Update-a-proximity-placement-group
	proximityPlacementGroupsClient, err := armcompute.NewProximityPlacementGroupsClient(subscriptionId, cred, nil)
	if err != nil {
		panic(err)
	}
	proximityPlacementGroupsClientCreateOrUpdateResponse, err := proximityPlacementGroupsClient.CreateOrUpdate(ctx, resourceGroupName, resourceName, armcompute.ProximityPlacementGroup{
		Location: to.Ptr(location),
		Properties: &armcompute.ProximityPlacementGroupProperties{
			ProximityPlacementGroupType: to.Ptr(armcompute.ProximityPlacementGroupTypeStandard),
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fakeScenarioVar = *proximityPlacementGroupsClientCreateOrUpdateResponse.ID

	// From step Delete-proximity_placement_group
	_, err = proximityPlacementGroupsClient.Delete(ctx, resourceGroupName, resourceName, nil)
	if err != nil {
		panic(err)
	}

	// From step Create_a_vm_with_Host_Encryption_using_encryptionAtHost_property
	virtualMachinesClient, err := armcompute.NewVirtualMachinesClient(subscriptionId, cred, nil)
	if err != nil {
		panic(err)
	}
	virtualMachinesClientCreateOrUpdateResponsePoller, err := virtualMachinesClient.BeginCreateOrUpdate(ctx, resourceGroupName, "myVM", armcompute.VirtualMachine{
		Location: to.Ptr(location),
		Plan: &armcompute.Plan{
			Name:      to.Ptr("signalrswaggertest6"),
			Product:   to.Ptr("windows-data-science-vm"),
			Publisher: to.Ptr("microsoft-ads"),
		},
		Properties: &armcompute.VirtualMachineProperties{
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardDS1V2),
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: to.Ptr("/subscriptions/" + subscriptionId + "/resourceGroups/" + resourceGroupName + "/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
						Properties: &armcompute.NetworkInterfaceReferenceProperties{
							Primary: to.Ptr(true),
						},
					}},
			},
			OSProfile: &armcompute.OSProfile{
				AdminPassword: to.Ptr("{your-password}"),
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("myVM"),
			},
			SecurityProfile: &armcompute.SecurityProfile{
				EncryptionAtHost: to.Ptr(true),
			},
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{
					Offer:     to.Ptr("windows-data-science-vm"),
					Publisher: to.Ptr(fakeScenarioVar),
					SKU:       to.Ptr("windows2016"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Name:         to.Ptr("myVMosdisk"),
					Caching:      to.Ptr(armcompute.CachingTypesReadOnly),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	_, err = virtualMachinesClientCreateOrUpdateResponsePoller.PollUntilDone(ctx, nil)
	if err != nil {
		panic(err)
	}
}

// Microsoft.SignalRService/DeleteOnly
func scenario1Sample() {
	// From step Delete_proximity_placement_group
	proximityPlacementGroupsClient, err := armcompute.NewProximityPlacementGroupsClient(subscriptionId, cred, nil)
	if err != nil {
		panic(err)
	}
	_, err = proximityPlacementGroupsClient.Delete(ctx, resourceGroupName, resourceName, nil)
	if err != nil {
		panic(err)
	}
}

func createResourceGroup() error {
	rand.Seed(time.Now().UnixNano())
	resourceGroupName = fmt.Sprintf("go-sdk-sample-%d", rand.Intn(1000))
	rgClient, err := armresources.NewResourceGroupsClient(subscriptionId, cred, nil)
	if err != nil {
		panic(err)
	}
	param := armresources.ResourceGroup{
		Location: to.Ptr(location),
	}
	_, err = rgClient.CreateOrUpdate(ctx, resourceGroupName, param, nil)
	if err != nil {
		panic(err)
	}
	return nil
}

func deleteResourceGroup() error {
	rgClient, err := armresources.NewResourceGroupsClient(subscriptionId, cred, nil)
	if err != nil {
		panic(err)
	}
	pollerResponse, err := rgClient.BeginDelete(ctx, resourceGroupName, nil)
	if err != nil {
		panic(err)
	}
	_, err = pollerResponse.PollUntilDone(ctx, nil)
	if err != nil {
		panic(err)
	}
	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func createDeployment(deploymentName string, deployment *armresources.Deployment) *armresources.DeploymentExtended {
	deployClient, err := armresources.NewDeploymentsClient(subscriptionId, cred, nil)
	if err != nil {
		panic(err)
	}
	poller, err := deployClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		deploymentName,
		*deployment,
		&armresources.DeploymentsClientBeginCreateOrUpdateOptions{},
	)
	if err != nil {
		panic(err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		panic(err)
	}
	return &res.DeploymentExtended
}

func generateAlphaNumericID(prefix string, length int) string {
	rand.Seed(time.Now().Unix())
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return prefix + string(b)
}
