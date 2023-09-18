package main

import "fmt"

type Entry struct {
	EventVersion string `json:"EventVersion"`
	EventSource  string `json:"EventSource"`
	Sns          struct {
		Message string `json:"Message"`
	} `json:"Sns"`
}

type SNS struct {
	Records []Entry `json:"Records"`
}

func main() {
	var message SNS
	var entry Entry

	entry.Sns.Message = "Hello from SNS!"
	entry.EventSource = "aws:sns"
	entry.EventVersion = "1.0"

	message.Records = make([]Entry, 1)
	message.Records[0] = entry

	fmt.Println(message)
}
