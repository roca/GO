package main

import (
	"fmt"
	"log"
	"math/rand"
)

// var count = 0

func generateIV(length int) ([]byte, error) {
	iv := make([]byte, length)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}
	return iv, nil
}

// don't touch below this line

func main() {
	rand.Seed(0)
	for i := 8; i < 17; i++ {
		iv, err := generateIV(i)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("%v-byte iv: %0*X, length of iv: %d\n", i, 2*i, iv, len(iv))
	}
}
