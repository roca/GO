package main

import "testing"

func Test_getOffsetChar(t *testing.T) {
	tests := []struct {
		name   string
		c      rune
		offset int
		want   string
	}{
		// 1
		{"a", 'a', 1, "b"},
		{"b", 'b', 1, "c"},
		{"x", 'x', 1, "y"},
		{"z", 'z', 1, "a"},
		// 5
		{"a", 'a', 5, "f"},
		{"b", 'b', 5, "g"},
		{"x", 'x', 5, "c"},
		{"z", 'z', 5, "e"},
		// 25
		{"A", 'A', 25, ""},
		// -5
		{"a", 'a', -5, "v"},
		{"b", 'b', -5, "w"},
		{"c", 'c', -5, "x"},
		{"d", 'd', -5, "y"},
		{"e", 'e', -5, "z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOffsetChar(tt.c, tt.offset); got != tt.want {
				t.Errorf("getOffsetChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crypt(t *testing.T) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	shifted := "defghijklmnopqrstuvwxyzabc"
	key := 3
	got := crypt(alphabet, key)
	if got != shifted {
		t.Errorf("crypt() = %v, want %v", got, shifted)
	}
}

func Test_decrypt(t *testing.T) {
	tests := []struct {
		name string
		text string
		key  int
	}{
		{"abcdefghi", "abcdefghi", 1},
		{"hello", "hello", 5},
		{"correcthorsebatterystaple", "correcthorsebatterystaple", 16},
		{"onetwothreefourfivesixseveneightnineten", "onetwothreefourfivesixseveneightnineten", 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted := crypt(tt.text, tt.key)
			decrypted := decrypt(encrypted, tt.key)
			if decrypted != tt.text {
				t.Errorf("decrypt() = %v, want %v", decrypted, tt.text)
			}
		})
	}
}

func Test_encrypt(t *testing.T) {
	tests := []struct {
		name string
		text string
		key  int
	}{
		{"abcdefghi", "abcdefghi", 1},
		{"hello", "hello", 5},
		{"correcthorsebatterystaple", "correcthorsebatterystaple", 16},
		{"onetwothreefourfivesixseveneightnineten", "onetwothreefourfivesixseveneightnineten", 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted := encrypt(tt.text, tt.key)
			decrypted := decrypt(encrypted, tt.key)
			if decrypted != tt.text {
				t.Errorf("decrypt() = %v, want %v", decrypted, tt.text)
			}
		})
	}
}
