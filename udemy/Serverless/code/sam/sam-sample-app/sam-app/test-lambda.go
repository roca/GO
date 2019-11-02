package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"

	"github.com/aws/aws-lambda-go/events"
)

func main() {

	c, err := rpc.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatalln(err)
	}

	file, _ := ioutil.ReadFile("event.json")

	data := events.APIGatewayProxyRequest{}

	_ = json.Unmarshal([]byte(file), &data)

	var result events.APIGatewayProxyResponse

	err = c.Call("Function.Invoke", &data, &result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Function.Invoke =", string(result.Body))

}
