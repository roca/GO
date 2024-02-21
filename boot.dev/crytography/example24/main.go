package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func genKeys() (pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey, err error) {
	// ?
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return publicKey, privateKey, nil
}

// don't touch below this line

func keysArePaired(pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey) bool {
	msg := "a test message"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		return false
	}

	return ecdsa.VerifyASN1(pubKey, hash[:], sig)
}

func test(i int) {
	fmt.Printf("Generating key pair %v...\n", i)
	pub, priv, err := genKeys()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Verifying key pair...")
	arePaired := keysArePaired(pub, priv)
	fmt.Printf("Keys are securely paired: %v\n", arePaired)
	fmt.Println("========")
}

func main() {
	for i := 1; i < 4; i++ {
		test(i)
	}
}
