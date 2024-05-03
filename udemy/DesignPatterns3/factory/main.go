package main

import (
	"factory/products"
	"fmt"
)

func main() {
	// Create a new product
	product := products.New("Laptop")

	// Print the product
	fmt.Printf("%+v\n", product)
}
