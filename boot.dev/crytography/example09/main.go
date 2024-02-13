package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	// ?
}

// don't touch below this line

func crypt(dat, key []byte) []byte {
	final := []byte{}
	for i, d := range dat {
		final = append(final, d^key[i])
	}
	return final
}

func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(num))
	if err != nil {
		fmt.Println("Error in intToBytes:", err)
		return nil
	}
	bs := buf.Bytes()
	if len(bs) > 3 {
		return bs[:3]
	}
	return bs
}

func test(encrypted []byte, decrypted string) {
	fmt.Printf("Encrypted: %x, decrypted: %s\n", []byte(encrypted), decrypted)
	fmt.Println("Starting brute force search...")
	key, err := findKey(encrypted, decrypted)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Key found: %x\n", key)
	fmt.Println("========")
}

func main() {
	test([]byte{0x1b, 0x2c, 0x3d}, "yes")
	test([]byte{0x2a, 0xff, 0xea}, "car")
	test([]byte{0x7d, 0x31, 0x32}, "she")
}
