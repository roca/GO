package main

import (
	"testing"
)

func Test_getHexString(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  string
	}{
		{"HexHello", []byte("Hello"), "48:65:6c:6c:6f"},
		{"HexWorld", []byte("World"), "57:6f:72:6c:64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHexString(tt.bytes); got != tt.want {
				t.Errorf("getHexString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBinaryString(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  string
	}{
		{"BinHello", []byte("Hello"), "01001000:01100101:01101100:01101100:01101111"},
		{"BinWorld", []byte("World"), "01010111:01101111:01110010:01101100:01100100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBinaryString(tt.bytes); got != tt.want {
				t.Errorf("getBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
