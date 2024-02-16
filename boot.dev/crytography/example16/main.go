package main

import "fmt"

func padWithZeros(block []byte, desiredSize int) []byte {
	return append(block, make([]byte, desiredSize-len(block))...)
}

// don't touch below this line

func test(block []byte, desiredSize int) {
	fmt.Printf("Padding %v for a total of %v bytes...\n",
		block,
		desiredSize,
	)
	padded := padWithZeros(block, desiredSize)
	fmt.Printf("Result: %v\n", padded)
	fmt.Println("========")
}

func main() {
	test([]byte{0xFF}, 4)
	test([]byte{0xFA, 0xBC}, 8)
	test([]byte{0x12, 0x34, 0x56}, 12)
	test([]byte{0xFA}, 16)
}
