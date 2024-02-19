package main

import (
	"bytes"
	"testing"
)

func Test_encrypt(t *testing.T) {
	tests := []struct {
		name      string
		key       []byte
		plaintext []byte
	}{
		{"test1", []byte("12344321"), []byte("Today I met my crush, what a hunk")},
		{"test2", []byte("p@$$w0rd"), []byte("I hope my boyfriend never finds out about this")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ciphertext, err := encrypt(tt.key, tt.plaintext)
			if err != nil {
				t.Error(err)
				return
			}
			decryptedText, err := decrypt(tt.key, ciphertext)
			if err != nil {
				t.Error(err)
				return
			}
			decryptedText = bytes.Trim(decryptedText, "\x00")
			if string(decryptedText) != string(tt.plaintext) {
				t.Errorf("Decrypted: '%v', want '%v'", string(decryptedText), string(tt.plaintext))
			}
		})
	}
}

func Test_padMsg(t *testing.T) {
	tests := []struct {
		name         string
		plaintext    []byte
		blockSize    int
		lengthWanted int
	}{
		{"test1", make([]byte, 9), 3, 12},
		{"test2", make([]byte, 15), 8, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := padMsg(tt.plaintext, tt.blockSize); len(got) != tt.lengthWanted {
				t.Errorf("padMsg() = %d, want %d", len(got), tt.lengthWanted)
			}
		})
	}
}
