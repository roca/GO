package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	var (
		instanceId string
		err        error
	)
	ctx := context.Background()
	if instanceId, err = createEC2(ctx, "us-east-1"); err != nil {
		fmt.Printf("createEC2 error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Instance id: %s\n", instanceId)

}

func createEC2(ctx context.Context, region string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region), config.WithSharedConfigProfile("go-aws-sdk"))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %s", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)

	_, err = ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
		KeyName: aws.String("go-aws-demo"),
	})

	if err != nil {
		return "", fmt.Errorf("CreateKeyPair error: %s", err)
	}

	imageOutput, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{"099720109477"},
	})

	if err != nil {
		return "", fmt.Errorf("DescribeImages error: %s", err)
	}
	if len(imageOutput.Images) == 0 {
		return "", fmt.Errorf("imageOutput.Images is of 0 length")
	}

	instance, err := ec2Client.RunInstances(ctx, &ec2.RunInstancesInput{
		ImageId: imageOutput.Images[0].ImageId,
		InstanceType: types.InstanceTypeT3Micro,
		KeyName: aws.String("go-aws-demo"),
		MinCount: aws.Int32(1),
		MaxCount: aws.Int32(1),

	})
	if err != nil {
		return "", fmt.Errorf("RunInstances error: %s", err)
	}
	if len(instance.Instances) == 0 {
		return "", fmt.Errorf("instance.Instances is of 0 length")
	}

	return *instance.Instances[0].InstanceId, nil
}
