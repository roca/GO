package main

import (
	"crypto/sha256"
	"fmt"
)

func splitInHalf(s string) (string, string) {
	firstHalf := s[:len(s)/2]
	secondHalf := s[len(firstHalf):]

	return firstHalf, secondHalf
}

func hmac(message, key string) string {
	keyFirstHalf, keySecondHalf := splitInHalf(key)

	h := sha256.New()
	h.Write([]byte(keySecondHalf + message))
	h.Write([]byte(keyFirstHalf + fmt.Sprintf("%x", h.Sum(nil))))

	return fmt.Sprintf("%x", h.Sum(nil))
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
