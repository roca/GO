package main

import (
	"fmt"
	"math/bits"
)

func hash(input []byte) [4]byte {
	final := [4]byte{}

	for i, b := range input {
		// do something with b
		rotated := bits.RotateLeft8(uint8(b), 3)
		shifted := rotated << 2
		final[i%len(final)] ^= shifted
	}
	return final
}

// don't touch below this line

func test(input string) {
	fmt.Printf("Hashing '%s'...\n", input)
	fmt.Printf("Output: %X\n", hash([]byte(input)))
	fmt.Println("========")
}

func main() {
	test("Example message")
	test("This is a slightly longer example to hash")
	test("This is a much longer example of some text to hash, maybe it's the opening paragraph of a blog post")
}
