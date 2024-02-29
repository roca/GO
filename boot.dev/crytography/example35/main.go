package main

import (
	"crypto/sha256"
	"fmt"
)

func hmac(message, key string) string {
	keyFirstHalf := key[:len(key)/2]
	keySecondHalf := key[len(keyFirstHalf):]

	innerDat := sha256.Sum256([]byte(keySecondHalf + message))
	allDat := []byte(keyFirstHalf)
	allDat = append(allDat, innerDat[:]...)

	sum := sha256.Sum256(allDat)
	result := fmt.Sprintf("%x", sum)
	return result
}

// don't touch below this line

func test(message, key string) {
	fmt.Printf("Calculating HMAC of '%v' with key '%v'...\n", message, key)
	checksum := hmac(message, key)
	fmt.Printf("HMAC: %v\n", checksum)
	fmt.Println("========")
}

func main() {
	test("I hope no one finds the Bitcoin keys I keep under my mailbox", "super_secret_password")
	test("No really, they're just written on a piece of paper", "correct horse battery staple")
	test("It's like a gazillion satoshis worth of BTC", "aFiveDoll@rWr3nch")
}
