package main

import "fmt"

func main() {

	m := [2][2]float64{}
	fmt.Println(m)

	for i, v1 := range m {
		for j, v2 := range v1 {
			fmt.Printf("(%d %d) %g ", i, j, v2)
		}
		fmt.Printf("\n")
	}

}
