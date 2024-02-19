package main

import "testing"

func Test_sBox(t *testing.T) {
	tests := []struct {
		name    string
		b       byte
		want    byte
		wantErr bool
	}{
		{"0", 0, 0b00, false},
		{"1", 1, 0b10, false},
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
