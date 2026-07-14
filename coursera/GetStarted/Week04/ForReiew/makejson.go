package main

import (
	"encoding/json"
	"fmt"
)

/* Write a program which prompts the user to first enter a name,
and then enter an address. Your program should create a map and
add the name and address to the map using the keys â€œnameâ€ and â€œaddressâ€,
respectively. Your program should use Marshal() to create a JSON object
from the map, and then your program should print the JSON object. */

func main() {
	var name string
	var address string

	type myJson struct {
		name    string
		address string
	}

	fmt.Print("Enter a name: ")
	fmt.Scan(&name)
	fmt.Print("Enter an address: ")
	fmt.Scan(&address)

	ourMap := make(map[string]string)
	ourMap["name"] = name
	ourMap["address"] = address

	barr, err := json.Marshal(ourMap)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(barr))

}
