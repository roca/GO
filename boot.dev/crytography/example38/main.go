package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func generateSalt(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func hashPassword(password, salt []byte) []byte {
	h := sha256.New()
	h.Write(append(password, salt...))
	return h.Sum(nil)
}

// don't touch below this line

func test(password1, password2 string, saltLen int) {
	defer fmt.Println("========")
	fmt.Printf("Hashing '%s' with salt length %v...\n", password1, saltLen)
	salt, err := generateSalt(saltLen)
	if err != nil {
		fmt.Printf("Error generating salt: %v", err)
		return
	}
	hashed := hashPassword([]byte(password1), salt)
	fmt.Println("Hash generated")

	fmt.Printf("Checking first hash against hash of '%v'...\n", password2)
	hashed2 := hashPassword([]byte(password2), salt)

	if string(hashed) == string(hashed2) {
		fmt.Println("Hashes match!")
	} else {
		fmt.Println("Hashes don't match!")
	}
}

func main() {
	test("samepass", "samepass", 16)
	test("passone", "passtwo", 24)
	test("correct horse battery staple", "correct horse battery staple", 32)
}
