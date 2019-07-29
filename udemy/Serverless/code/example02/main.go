package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

// Event ...
type Event struct {
}

var mesages = []string{
	"Hello World",
	"Hello Serverless",
	"It's a great day today",
	"Yay, I'm learning something new today",
	"On cloud nine!",
	"Over the moon",
	"Shooting for the stars!",
	"On top of the World",
	"World at my feet!",
	"Doing everthing I love!",
}

func handler(event Event) (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	message := mesages[r.Intn(9)]
	return fmt.Sprintf(message), nil
}

func main() {
	lambda.Start(handler)
}
