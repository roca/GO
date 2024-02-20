package main

import (
	"bytes"
	"testing"
)

func Test_decrypt(t *testing.T) {
	tests := []struct {
		name      string
		key       []byte
		plaintext []byte
		nounce    []byte
	}{
		{
			"test1",
			[]byte("d00c5215-60f6-4ac4-9648-532b5dad"),
			[]byte("Today I met my crush, what a hunk"),
			generateNonce(12),
		},
		{
			"test2", 
			[]byte("db50ecaaa-23ed-43eb-9f8b-6fc5931"), 
			[]byte("I knew it, Becky has been cheating this whole time!"),
			generateNonce(12),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ciphertext, err := encrypt(tt.key, tt.plaintext,tt.nounce)
			if err != nil {
				t.Error(err)
				return
			}
			decryptedText, err := decrypt(tt.key, ciphertext, tt.nounce)
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
