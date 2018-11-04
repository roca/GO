// File name: ...\s05\01_maps_hashtables\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// This example is only to demonstrate how hashtables work.
// The logic to make hashkeys is not optimized.
// This solution only works for up to 3 characters.
func main() {
	values := []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
	// values := []string{"to", "to", "top", "ton", "tom"}
	factor := []int{100, 10, 1}

	// 65x100 + 66x10 + 67x1 = 7227
	hashKey := 0
	for v := range values {
		bytes := []byte(values[v])
		f := 0
		hashKey = 0
		for i := range bytes {
			fmt.Print(bytes[i], " ")
			hashKey += int(bytes[i]) * factor[f]
			f++
		}
		fmt.Printf(" (hashKey: %d) \n", hashKey)
	}
}
