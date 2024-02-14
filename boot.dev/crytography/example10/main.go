package main

import (
	"fmt"
	"strings"
)

func encrypt(plaintext string, key int) string {
	// ?
	return ""
}

func decrypt(ciphertext string, key int) string {
	// ?
	return ""
}

func crypt(text string, key int) string {
	// ?
	return ""
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	// ?
	i := strings.IndexRune(alphabet, c)
	return string(alphabet[(i+offset)])
}

// don't touch below this line

func test(plaintext string, key int) {
	fmt.Printf("Encrypting %v with key %v\n", plaintext, key)
	encrypted := encrypt(plaintext, key)
	fmt.Printf("Encrypted text: %v\n", encrypted)
	decrypted := decrypt(encrypted, key)
	fmt.Printf("Decrypted text: %v\n", decrypted)
	fmt.Println("========")
}

func main() {
	test("abcdefghi", 1)
	test("hello", 5)
	test("correcthorsebatterystaple", 16)
	test("onetwothreefourfivesixseveneightnineten", 25)
}
