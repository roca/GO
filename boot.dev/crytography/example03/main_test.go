package main

import "testing"

func Test_keyToCipher(T *testing.T) {
	const symmetricKey = "thisIsMySecretKeyIHopeNoOneFinds"
	cipher, err := keyToCipher(symmetricKey)
	if err != nil {
		T.Errorf("Expected nil, got %v", err)
	}
	if cipher.BlockSize() != 16 {
		T.Errorf("Expected 16, got %v", cipher.BlockSize())
	}
}
