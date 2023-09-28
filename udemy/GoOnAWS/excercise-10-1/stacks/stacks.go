package stacks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

var Client *cloudformation.Client

type Stack struct {
	Name   string
	Status string
}

func (s Stack) String() string {
	return s.Name + " " + s.Status
}

type Resource struct {
	LogicalID string
	Status    string
}

func (r Resource) String() string {
	return r.LogicalID + " " + r.Status
}

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = cloudformation.NewFromConfig(cfg)

}

func ListStacks(client *cloudformation.Client) ([]Stack, error) {
	stacks := []Stack{}

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

func StackResources(client *cloudformation.Client, stackName *string) ([]Resource, error) {
	resources := []Resource{}

	input := &cloudformation.DescribeStackResourcesInput{
		StackName: stackName,
	}

	resp, err := client.DescribeStackResources(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	for _, resource := range resp.StackResources {
		resources = append(resources, Resource{
			LogicalID: *resource.LogicalResourceId,
			Status:    string(resource.ResourceStatus),
		})
	}

	return resources, nil
}
