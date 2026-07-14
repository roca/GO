package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func (h *hasher) Write(s string) (int,error){
	return h.hash.Write([]byte(s))
}

func (h *hasher) GetHex() string {
	return fmt.Sprintf("%x",h.hash.Sum(nil))
}

func newHasher() *hasher {
	h := hasher{}
	h.hash = sha256.New()
	return &h
}

// don't touch below this line

func test(passwords []string) {
	fmt.Printf("Hashing vault of length %v...\n", len(passwords))
	h := newHasher()
	for _, password := range passwords {
		h.Write(password)
		fmt.Printf("Adding '%v' to vault hash...\n", password)
	}
	fmt.Printf("Vault hash: %v\n", h.GetHex())
	fmt.Println("========")
}

func main() {
	test([]string{"password1", "password2", "password3"})
	test([]string{"abercromni3", "f1tch", "123456", "abcdefg1234"})
	test([]string{"IHeartNanciedrake", "m7B1rthd@y"})
}
