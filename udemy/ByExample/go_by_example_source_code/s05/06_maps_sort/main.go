// File name: ...\s05\06_maps_sort\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"sort"
)

// ASSIGNEMENT: Maps are unordered. Write a program to use
// a "map[string]float64" data map and sort it.
func main() {
	sal := map[string]float64{
		"Blake":  60000.00,
		"Parker": 120000.50,
		"Dakota": 93000.00,
	}

	names := make([]string, 0, len(sal))
	for n := range sal {
		names = append(names, n)
	}
	fmt.Println(names)

	sort.Strings(names)
	fmt.Println(names)

	for _, n := range names {
		fmt.Printf("%s\t%.2f\n", n, sal[n])
	}

}
