package main

import "fmt"

func main() {
	var varInt = 42
	var varString = "mega"
	var vrFloat = 3.14
	varBool := true

	fmt.Printf("Value Int : %v\n", varInt)
	fmt.Printf("Type Int : %T\n", varInt)

	fmt.Printf("Value Int : %v\n", varString)
	fmt.Printf("Type Int : %T\n", varString)

	fmt.Printf("Value Int : %v\n", vrFloat)
	fmt.Printf("Type Int : %T\n", vrFloat)

	fmt.Printf("Value Int : %v\n", varBool)
	fmt.Printf("Type Int : %T\n", varBool)
}
