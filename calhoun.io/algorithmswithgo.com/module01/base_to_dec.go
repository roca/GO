package module01

import "strings"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	var d int
	r := []rune(Reverse(strings.ToUpper(value)))
	for i, v := range r {
		n := int(v)
		if n >= 65 && n <= 90 {
			n -= 55
		} else if n >= 48 && n <= 57 {
			n -= 48
		}
		d += n * pow(base, i)
	}
	return int(d)
}

func pow(n, exponent int) int {
	opr := 1
	for i := exponent; i > 0; i-- {
		opr *= n
	}
	return opr
}
