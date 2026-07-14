package main

import "fmt"

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	var xor [4]byte

	for i, v := range masterKey {
		xor[i] = v ^ byte(roundNumber)
	}
	return xor
}

// don't touch below this line

func test(masterKey [4]byte) {
	fmt.Printf("Deriving round keys from master: %X...\n", masterKey)
	for i := 1; i < 9; i++ {
		roundKey := deriveRoundKey(masterKey, i)
		fmt.Printf(" - Round key %v: %X\n", i, roundKey)
	}
	fmt.Println("========")
}

func main() {
	test([4]byte{0xAA, 0xFF, 0x11, 0xBC})
	test([4]byte{0xEB, 0xCD, 0x13, 0xFC})
}
