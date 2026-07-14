package main

import (
	"bytes"
	"testing"
)

func Test_getHexBytes(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []byte
	}{
		{"HexnHello", "48:65:6c:6c:6f", []byte("Hello")},
		{"HexWorld", "57:6f:72:6c:64", []byte("World")},
		{"HexPassword", "50:61:73:73:77:6f:72:64", []byte("Password")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getHexBytes(tt.text)
			if err != nil {
				t.Errorf("getHexBytes() error = %v", err)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("getHexBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
