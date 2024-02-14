package main

import "testing"

func Test_getOffsetChar(t *testing.T) {
	tests := []struct {
		name   string
		c      rune
		offset int
		want   string
	}{
		{"a", 'a', 1, "b"},
		{"a", 'a', 5, "f"},
		{"A", 'A', 25, "z"},
		//{"a", 'a', -5, "v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOffsetChar(tt.c, tt.offset); got != tt.want {
				t.Errorf("getOffsetChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
