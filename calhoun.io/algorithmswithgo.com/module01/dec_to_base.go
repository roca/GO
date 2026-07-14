package module01

import (
	"strings"
)

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
	var sb strings.Builder
	for i := dec; i > 0; i /= base {
		m := i % base
		if m >= 0 && m < 10 {
			sb.WriteRune(rune(m + 48))
		} else {
			sb.WriteRune(rune(m + 55))
		}
	}
	return Reverse(sb.String())
}
