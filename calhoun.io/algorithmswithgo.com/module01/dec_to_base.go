package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	s := ""
	for i := dec; i > 0; i /= base {
		m := i % base
		if m < 10 {
			s = fmt.Sprintf("%d", m) + s
		} else {
			s = fmt.Sprintf("%c", rune(m+55)) + s
		}
	}
	return s
}
