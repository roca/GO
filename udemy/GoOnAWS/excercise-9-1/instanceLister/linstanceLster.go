package instanceLister

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var Client *ec2.Client

type Instances struct {
	Name  string
	State string
}

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = ec2.NewFromConfig(cfg)

}

func ListInstances(client *ec2.Client) ([]*Instances, error) {
	instances := []*Instances{}

	parms := &ec2.DescribeInstancesInput{
		MaxResults: aws.Int32(10),
	}

	paginator := ec2.NewDescribeInstancesPaginator(client, parms)
	//pageNum := 1
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		//fmt.Println("Page: ", pageNum)

		//count := len(output.Reservations)
		//fmt.Println("Instances: ", count)
		for _, reservation := range output.Reservations {
			for _, instance := range reservation.Instances {
				newInstance := &Instances{
					Name:  *instance.InstanceId,
					State: string(instance.State.Name),
				}
				instances = append(instances, newInstance)
			}
		}

		// pageNum += 1
	}

	return instances, nil
}
