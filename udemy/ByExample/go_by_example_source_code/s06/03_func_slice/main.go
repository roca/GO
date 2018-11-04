// File name: ...\s06\03_func_slice\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	scores1 := []float32{91, 82, 99}
	fmt.Printf("average: %.2f\n", avg(scores1))

	scores2 := []float32{72, 81, 78, 91, 68}
	fmt.Printf("average: %.2f\n", avg(scores2))
}

func avg(scores []float32) float32 {
	var total float32
	for _, score := range scores {
		total += score
	}
	return total / float32(len(scores))
}
