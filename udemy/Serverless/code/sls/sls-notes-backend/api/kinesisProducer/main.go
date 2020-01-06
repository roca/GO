package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var sess *session.Session
var svc *kinesis.Kinesis
var streamName *string

func init() {

	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	svc = kinesis.New(sess)
	streamName = aws.String("ServerlessNotesStream")
}

func main() {

	out, err := svc.CreateStream(&kinesis.CreateStreamInput{
		ShardCount: aws.Int64(1),
		StreamName: streamName,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)

	if err := svc.WaitUntilStreamExists(&kinesis.DescribeStreamInput{StreamName: streamName}); err != nil {
		panic(err)
	}

	streams, err := svc.DescribeStream(&kinesis.DescribeStreamInput{StreamName: streamName})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", streams)

}
