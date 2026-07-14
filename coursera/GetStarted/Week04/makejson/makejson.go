/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Scanf wont work if your input has spaces :)
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter a name")
	name, _ := consoleReader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Println("Please enter a address")
	address, _ := consoleReader.ReadString('\n')
	address = strings.TrimSuffix(address, "\n")

	myMap := map[string]string{
		"name":    name,
		"address": address,
	}

	myJSONString, err := ReturnJSONString(myMap)
	if err != nil {
		log.Fatal("Could not covert this object to JSONi string")
	}

	fmt.Println(myJSONString)

}

// ReturnJSONString : converts a GO object to JSON string
func ReturnJSONString(o interface{}) (string, error) {
	jsonObject, err := json.Marshal(o)
	return string(jsonObject), err
}
