/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	var name, address string

	fmt.Println("Please enter a name")
	fmt.Scan(&name)

	fmt.Println("Please enter a address")
	fmt.Scan(&address)

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
	if err != nil {
		log.Fatal("Could not covert this object to JSON")
	}

	return string(jsonObject), err

}
