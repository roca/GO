package main

import "testing"

func Test_baseChar(t *testing.T) {
	tests := []struct {
		name string
		bits byte
		want string
	}{
		{"A", 0b000, "A"},
		{"B", 0b001, "B"},
		{"C", 0b010, "C"},
		{"D", 0b011, "D"},
		{"E", 0b100, "E"},
		{"F", 0b101, "F"},
		{"G", 0b110, "G"},
		{"H", 0b111, "H"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base8Char(tt.bits); got != tt.want {
				t.Errorf("base8Char() = %v, want %v", got, tt.want)
			}
		})
	}
	// 8 and above
	for i := 8; i < 256; i++ {
		if got := base8Char(byte(i)); got != "" {
			t.Errorf("base8Char() = %v, want empty string", got)
		}
	}
}
