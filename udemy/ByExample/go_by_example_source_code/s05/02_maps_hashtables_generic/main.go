// File name: ...\s05\02_maps_hashtables_generic\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math"
)

// ASSIGNMENT - improve the previous example by making it more generic.
func main() {
	values := []string{"tea", "tea", "eat", "eta", "teal", "teatop", "teaparty", "tear"}

	for v := range values {
		hashKey := hashK(values[v])
		fmt.Printf(" %-10s (hashKey: %10d) %v\n", values[v], hashKey, []byte(values[v]))
	}
}

func hashK(s string) int {
	result := 0
	idx := 0
	for i := len(s) - 1; i > -1; i-- {
		result += int(math.Pow10(i)) * int(s[idx])
		idx++
	}
	return result
}
