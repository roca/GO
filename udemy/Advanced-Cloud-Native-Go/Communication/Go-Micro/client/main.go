package main

import (
	"fmt"
	"time"

	micro "github.com/micro/micro/v3/service"
	"golang.org/x/net/context"
	proto "udemy.com/proto"
)

// The Greeter API.
type Greeter struct{}

// Hello is a Greeter API method.
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func callEvery(d time.Duration, greeter proto.GreeterService, f func(time.Time, proto.GreeterService)) {
	for x := range time.Tick(d) {
		f(x, greeter)
	}
}

func hello(t time.Time, greeter proto.GreeterService) {
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Leander, calling at " + t.String()})
	if err != nil {
		if err.Error() == "hystrix: timeout" {
			fmt.Printf("%s. Insert fallback logic here.\n", err.Error())
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	// Print response
	fmt.Printf("%s\n", rsp.Greeting)
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.New(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings.
	service.Init()

	// Create new greeter client and call hello
	greeter := proto.NewGreeterService("greeter", service.Client())
	callEvery(3*time.Second, greeter, hello)
}
