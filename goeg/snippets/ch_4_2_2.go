// ch_4_2_2
package main

import (
	"fmt"
)

type Product struct {
	name  string
	price float64
}

//func (product Product) String() string {
//	return fmt.Sprintf("%s (%.2f)", product.name, product.price)

//}

func main() {

	products := []*Product{{"Spanner", 3.99}, {"Wrench", 2.49}, {"Screwdriver", 1.99}}
	fmt.Println(products)
	for _, product := range products {
		product.price += 0.50
	}
	fmt.Println(products)
}
