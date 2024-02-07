package main

import "testing"

func Test_encrypt_decrypt(t *testing.T) {
	type args struct {
		plainText string
		key       string
		iv        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{}

	_ = tests

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ""
			if got != tt.want {
				t.Errorf("encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
