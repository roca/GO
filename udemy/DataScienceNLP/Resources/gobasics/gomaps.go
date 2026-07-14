package main

import "fmt"

func main() {

	// Map: key value pair collection
	// Dictionary/ Associative Collection
	// make : create slice,map ,channel

	// Method: Creating A Map
	// map[key]value

	var mydict map[string]int
	mydict = make(map[string]int)

	// Assign Values
	mydict["age"] = 24

	fmt.Println(mydict["age"])

	// Method 2: make and shorthand
	mydict2 := make(map[string]string)
	mydict2["name"] = "Jesse"
	mydict2["location"] = "London"

	fmt.Println(mydict2)
	// Select By Key
	fmt.Println("Location:", mydict2["location"])

	// Check if key is in a map/dictionary
	elem, ok := mydict["age"]
	fmt.Println(elem, ok)

	// Loop Through via range
	for key, value := range mydict2 {
		fmt.Println(key, "=", value)
	}

	// Get only values using blank identifier
	for _, value := range mydict2 {
		fmt.Println(value)
	}

	// Delete
	delete(mydict2, "name")
	fmt.Println(mydict2)

}
