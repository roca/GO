package main


func Test_encrypt(t *testing.T) {
	tests := []struct {
		name     string
		key      []byte
		plaintext []byte
	} {
		{"test1", []byte("1234321"), []byte("Today I met my crush, what a hunk")},
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
	