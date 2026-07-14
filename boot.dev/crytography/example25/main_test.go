package main

import "testing"

func Test_encrypt(t *testing.T) {

	tests := []struct {
		name      string
		plaintext []byte
		wantErr   bool
	}{
		{"test1", []byte("Hey Darling, don't come over tonight, I'm out with my people"), false},
		{"test2", []byte("Yes, ten million in cash. No, every penny better be accounted for"), false},
		{"test3", []byte("Do you know what would happen if I suddenly decided to stop going into work? A business big enough that it could be listed on the NASDAQ goes belly up. Disappears! It ceases to exist without me. No, you clearly don't know who you're talking to, so let me clue you in. I am not in danger, Skyler. I am the danger. A guy opens his door and gets shot and you think that of me? No. I am the one who knocks!"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pubKey, privKey, err := genKeys()
			if err != nil {
				t.Errorf("genKeys() error = %v", err)
				return
			}
			ciphertext, err := encrypt(pubKey, tt.plaintext)
			if (err != nil) != tt.wantErr {
				t.Errorf("encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				plaintext, err := decrypt(privKey, ciphertext)
				if (err != nil) != tt.wantErr {
					t.Errorf("decrypt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if string(plaintext) != string(tt.plaintext) {
					t.Errorf("decrypt() \ngot = %v, \nwant %v", string(plaintext), string(tt.plaintext))
					return
				}
			}
		})
	}

}
