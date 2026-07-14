// File name: ...\s02\15_min_max\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%30d %30d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("%30d %30d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("%30d %30d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("%30d %30d\n", math.MinInt64, math.MaxInt64)

	fmt.Println()
	fmt.Printf("%30d %30d\n", 0, math.MaxUint8)
	fmt.Printf("%30d %30d\n", 0, math.MaxUint16)
	fmt.Printf("%30d %30d\n", 0, math.MaxUint32)

	// compiler error- constant 18446744073709551615 overflows int
	// fmt.Printf("%30d %30d\n", 0, math.MaxUint64)
	// math.MaxUint64 is a constant, not an int64.
	fmt.Printf("%30d %30d\n", 0, uint64(math.MaxUint64))

	fmt.Println()
	fmt.Printf("%61v\n", math.MaxFloat32)
	fmt.Printf("%61v\n", math.MaxFloat64)
}
