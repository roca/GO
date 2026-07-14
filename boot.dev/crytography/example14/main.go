package main

import (
	"errors"
	"fmt"
)

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result)
	i := 0
	for {
		textChar, textOk := <-textCh
		if !textOk {
			return
		}
		keyChar, keyOk := <-keyCh
		if !keyOk {
			return
		}
		result <- textChar ^ keyChar
		i++
		fmt.Println("Crypted byte:", i)
	}
}

// don't touch below this line

func encrypt(plaintext, key []byte) ([]byte, error) {
	if len(plaintext) != len(key) {
		return nil, errors.New("plaintext and key must be the same length")
	}

	plaintextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	go func() {
		defer close(plaintextCh)
		for _, v := range plaintext {
			plaintextCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(plaintextCh, keyCh, result)

	res := []byte{}
	for v := range result {
		res = append(res, v)
	}
	return res, nil
}

func decrypt(ciphertext, key []byte) ([]byte, error) {
	if len(ciphertext) != len(key) {
		return nil, errors.New("ciphertext and key must be the same length")
	}

	ciphertextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	go func() {
		defer close(ciphertextCh)
		for _, v := range ciphertext {
			ciphertextCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(ciphertextCh, keyCh, result)

	res := []byte{}
	for v := range result {
		res = append(res, v)
	}
	return res, nil
}

func test(plaintext, key []byte) {
	fmt.Printf("Encrypting '%s' using key '%s'\n", string(plaintext), string(key))
	ciphertext, err := encrypt(plaintext, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypted ciphertext bytes: %v\n", ciphertext)
	decrypted, err := decrypt(ciphertext, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decrypted message: %v\n", string(decrypted))
	fmt.Println("========")
}

func main() {
	test([]byte("Shazam"), []byte("Sk7p13"))
	test([]byte("I'm lovin it"), []byte("mysecurepass"))
}
