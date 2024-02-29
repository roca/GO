package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	}
	return false
}

// don't touch below this line

func test(password1, password2 string) {
	defer fmt.Println("========")
	fmt.Printf("Hashing '%s'...\n", password1)
	hashed, err := hashPassword(password1)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		return
	}
	fmt.Printf("Bcrypt output generated with len: %v\n", len(hashed))
	match := checkPasswordHash(password2, hashed)
	fmt.Printf("%v has a matching hash: %v\n", password2, match)
}

func main() {
	test("thisIsAPassword", "thisIsAPassword")
	test("thisIsAPassword", "thisIsAnotherPassword")
	test("corr3ct h0rse", "corr3ct h0rse")
}
