package module01

import (
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	var d float64
	r := []rune(value)
	for i := len(r); i > 0; i-- {
		n := int(r[i-1])
		if n >= 65 && n <= 90 {
			n -= 55
		} else if n >= 48 && n <= 57 {
			n -= 48
		}
		d += float64(n) * math.Pow(float64(base), float64(len(r)-i))
	}
	return int(d)
}
