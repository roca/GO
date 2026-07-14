package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func encrypt(pubKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		pubKey,
		msg, nil,
	)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// don't touch below this line

func decrypt(privKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privKey,
		ciphertext,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func genKeys() (pubKey *rsa.PublicKey, privKey *rsa.PrivateKey, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return &privateKey.PublicKey, privateKey, nil
}

func test(pubKey *rsa.PublicKey, privKey *rsa.PrivateKey, msg string) {
	defer fmt.Println("========")
	fmt.Printf("Encrypting message: '%v'\n", msg)
	ciphertext, err := encrypt(pubKey, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Ciphertext created with length %v\n", len(ciphertext))
	plaintext, err := decrypt(privKey, ciphertext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Plaintext: %v\n", string(plaintext))
}

func main() {
	pub, priv, err := genKeys()
	if err != nil {
		fmt.Println(err)
		return
	}
	test(pub, priv, "Hey Darling, don't come over tonight, I'm out with my people")
	test(pub, priv, "Yes, ten million in cash. No, every penny better be accounted for")
	test(pub, priv, "Do you know what would happen if I suddenly decided to stop going into work? A business big enough that it could be listed on the NASDAQ goes belly up. Disappears! It ceases to exist without me. No, you clearly don't know who you're talking to, so let me clue you in. I am not in danger, Skyler. I am the danger. A guy opens his door and gets shot and you think that of me? No. I am the one who knocks!")
}
