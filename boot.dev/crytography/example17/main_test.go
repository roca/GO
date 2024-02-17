package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_deriveRoundKey(t *testing.T) {
	key := [4]byte{0xAA, 0xFF, 0x11, 0xBC}

	tests := []struct {
		name        string
		roundNumber int
		want        [4]byte
	}{
		{"Round key 1", 1, [4]byte{0xAB, 0xFE, 0x10, 0xBD}},
		{"Round key 5", 5, [4]byte{0xAF, 0xFA, 0x14, 0xB9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deriveRoundKey(key, tt.roundNumber)
			if len(got) != len(tt.want) {
				t.Errorf("Wrong lengths")
				return
			}
			if got != tt.want {
				t.Errorf("deriveRoundKey(%v, %d) = %v, want %v, roundNumber in byte: %08b", toString(key), tt.roundNumber, toString(got), toString(tt.want), byte(tt.roundNumber))
			}
		})
	}
}

func toString(bytes [4]byte) string {
	var s []string
	for _, v := range bytes {
		s = append(s, fmt.Sprintf("%08b", v))
	}
	return strings.Join(s, ":")
}
