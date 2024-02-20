package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"math/rand"
)

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err = aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// don't touch below this line

func encrypt(key, plaintext, nonce []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func test(key, plaintext, nonce []byte) {
	ciphertext, err := encrypt(key, plaintext, nonce)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypting '%v' with key '%v'...\n", string(plaintext), string(key))
	decryptedText, err := decrypt(key, ciphertext, nonce)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decrypted: '%v'\n", string(decryptedText))
	fmt.Println("========")
}

func generateNonce(length int) []byte {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	return randomBytes
}

func main() {
	rand.NewSource(0)
	test(
		[]byte("d00c5215-60f6-4ac4-9648-532b5dad"),
		[]byte("I wonder what he's thinking about me??"),
		generateNonce(12),
	)
	test(
		[]byte("db50ecaaa-23ed-43eb-9f8b-6fc5931"),
		[]byte("I knew it, Becky has been cheating this whole time!"),
		generateNonce(12),
	)
}
