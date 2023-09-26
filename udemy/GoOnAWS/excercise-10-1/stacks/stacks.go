package stacks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

var Client *cloudformation.Client

type Stack struct {
	Name   string
	Status string
}

type Stacks []Stack

type Resource struct {
	LogicalID string
	Status string
}

type Resources []Resource

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = cloudformation.NewFromConfig(cfg)

}

func ListStacks(client *cloudformation.Client) (Stacks, error) {
	stacks := Stacks{}

	input := &cloudformation.DescribeStacksInput{}
	resp, _ := client.DescribeStacks(context.TODO(), input)

	for _, stack := range resp.Stacks {
		stacks = append(stacks, Stack{
			Name:   *stack.StackName,
			Status: string(stack.StackStatus),
		})
	}

	return stacks, nil
}

func ListStackResources(client *cloudformation.Client, stackName string) (Resources, error) {
	resources := Resources{}

	input := &cloudformation.DescribeStackResourcesInput{
		StackName: aws.String(stackName),
	}

	resp, err := client.DescribeStackResources(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	for _, resource := range resp.StackResources {
		resources = append(resources, Resource{
			LogicalID: *resource.LogicalResourceId,
			Status: string(resource.ResourceStatus),
		})
	}

	return resources, nil
}
