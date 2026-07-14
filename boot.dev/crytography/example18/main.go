package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/bits"
)

func feistel(msg []byte, roundKeys [][]byte) []byte {
	leftSide := msg[:len(msg)/2]
	rightSide := msg[len(msg)/2:]
	for _, key := range roundKeys {
		newRight := xor(leftSide, hash(rightSide, key, len(leftSide)))
		leftSide = rightSide
		rightSide = newRight
	}
	return append(rightSide, leftSide...)
}

// don't touch below this line

func test(msg []byte, key []byte, rounds int) {
	roundKeys := [][]byte{}
	for i := 0; i < rounds; i++ {
		ui := binary.BigEndian.Uint32(key)
		rotated := bits.RotateLeft32(uint32(ui), i)
		finalRound := make([]byte, len(key))
		binary.LittleEndian.PutUint32(finalRound, uint32(rotated))
		roundKeys = append(roundKeys, finalRound)
	}

	fmt.Printf("Encrypting '%v' with %v round keys...\n", string(msg), rounds)
	encrypted := feistel(msg, roundKeys)
	decrypted := feistel(encrypted, reverse(roundKeys))
	fmt.Printf("Decrypted: '%v'\n", string(decrypted))
	fmt.Println("========")
}
func main() {
	test(
		[]byte("General Kenobi!!!!"),
		[]byte("thesecret"),
		8,
	)
	test(
		[]byte("Hello there!"),
		[]byte("@n@kiN"),
		16,
	)
}

func reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func xor(lhs, rhs []byte) []byte {
	res := []byte{}
	for i := range lhs {
		res = append(res, lhs[i]^rhs[i])
	}
	return res
}

// outputLength should be equal to the key length
// when used in feistel so that the XOR operates on
// inputs of the same size
func hash(first, second []byte, outputLength int) []byte {
	h := sha256.New()
	h.Write(append(first, first...))
	return h.Sum(nil)[:outputLength]
}
