package main

import (
	"math/big"
	"testing"
)

func Test_Decrypt(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		keySize int
	}{
		{"test1", "I hid the cash under the sink", 512},
		{"test2", "Don't you think they will look there??", 512},
		{"test3", "They'll look at everything but the kitchen sink", 1024},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msgBytes := []byte(tt.msg)
			p, q := generatePrivateNums(tt.keySize)
			n := getN(p, q)
			phi := getPhi(p, q)
			e := getE(phi)
			plaintext := big.NewInt(0)
			plaintext.SetBytes(msgBytes)

			ciphertext := encrypt(plaintext, e, n)

			d := getD(e, phi)
			decrypted := decrypt(ciphertext, d, n)

			if string(decrypted.Bytes()) != tt.msg {
				t.Errorf("Decrypted message: got %s, want %s", string(decrypted.Bytes()), tt.msg)
			}	
		})
	}
}
