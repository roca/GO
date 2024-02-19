package main

import "testing"

func Test_sBox(t *testing.T) {
	tests := []struct {
		name    string
		b       byte
		want    byte
		wantErr bool
	}{
		{"0", 0b0000, 0b00, false},
		{"1", 0b0001, 0b10, false},
		{"15", 0b1111, 0b00, false},
		{"6",0b0110, 0b11, false},
		{"16", 0b10000, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sBox(tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("sBox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
