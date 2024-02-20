package main

import (
	"errors"
	"fmt"
)

func sBox(b byte) (byte, error) {
	if b > 0b1111 {
		return 0, errors.New("invalid input")
	}
	box := [4][4]byte{
		{0b00, 0b10, 0b01, 0b11},
		{0b10, 0b00, 0b11, 0b01},
		{0b01, 0b11, 0b00, 0b10},
		{0b11, 0b01, 0b10, 0b00},
	}
	// 0 2 1 3
	// 2 0 3 1
	// 1 3 0 2
	// 3 1 2 0

	row := (b & 0b1100) >> 2
	col := b & 0b0011
	return box[row][col], nil
}

// don't touch below this line

func main() {
	for i := 0; i <= 16; i++ {
		b := byte(i)
		subbed, err := sBox(b)
		if err != nil {
			fmt.Printf("Error with input %04b: %v\n", i, err)
			continue
		}
		fmt.Printf("%04b -> %02b\n", i, subbed)
	}
}
