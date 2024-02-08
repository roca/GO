package main

import "testing"

func Test_encrypt_decrypt(t *testing.T) {
	const masterKey = "kjhgfdsaqwertyuioplkjhgfdsaqwert"
	const iv = "1234567812345678"
	type args struct {
		plainText string
		key       string
		iv        string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{"k33pThisPasswordSafe", masterKey, iv}},
		{"test2", args{"12345", masterKey, iv}},
		{"test3", args{"thePasswordOnMyLuggage", masterKey, iv}},
		{"test4", args{"pizza_the_HUt", masterKey, iv}},
	}

	_ = tests

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decrypt(encrypt(tt.args.plainText, masterKey, iv), masterKey, iv)
			want := tt.args.plainText
			if got != tt.args.plainText {
				t.Errorf("encrypt() = %v, want %v", got, want)
			}
		})
	}
}
