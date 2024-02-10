package main

import (
	"fmt"
	"log"
	"crypto/rand"
)

func generateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x",string(bytes)), nil
}

// don't touch below this line

func main() {
	for i := 16; i < 33; i++ {
		key, err := generateRandomKey(i)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v-byte key: %v\n", i, key)
	}
}
