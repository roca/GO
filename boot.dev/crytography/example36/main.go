package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"strings"
)

func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write([]byte(message))
	hash := h.Sum(nil)
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hash)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%x", message, signature), nil
}

// don't touch below this line

func verifyECDSAMessage(token string, publicKey *ecdsa.PublicKey) error {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return errors.New("invalid token sections")
	}
	sig, err := hex.DecodeString(parts[1])
	if err != nil {
		return err
	}
	hash := sha256.Sum256([]byte(parts[0]))

	valid := ecdsa.VerifyASN1(publicKey, hash[:], sig)
	if !valid {
		return errors.New("invalid signature")
	}
	return nil
}

func test(message string) {
	defer fmt.Println("========")
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Creating token for: '%v'...\n", message)
	token, err := createECDSAMessage(message, privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Verifying token...")
	err = verifyECDSAMessage(token, &privateKey.PublicKey)
	if err != nil {
		fmt.Printf("Token is invalid: %v\n", err)
		return
	}
	fmt.Println("Token is valid!")
}

func main() {
	test("userid:2f9c584e-5d25-4516-a0ed-ddfa6e152006")
	test("userid:0e803af6-292f-4432-a285-84a7591e25a8")
	test("userid:f77e36d6-0edc-44ef-964e-af4a5b1ebd5f")
}
