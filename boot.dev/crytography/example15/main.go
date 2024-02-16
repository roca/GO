package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
)

const (
	typeAES = iota
	typeDES
)

func getBlockSize(keyLen, cipherType int) (int, error) {
	var cipherBlock cipher.Block
	var err error
	bytes := make([]byte, keyLen)

	switch cipherType {
	case typeAES:
		cipherBlock, err = aes.NewCipher(bytes)
	case typeDES:
		cipherBlock, err = des.NewCipher(bytes)
	default:
		return 0, errors.New("invalid cipher type")
	}
	if err != nil {
		return 0, err
	}
	return cipherBlock.BlockSize(), nil
}

// don't touch below this line

func test(keyLen, cipherType int) {
	fmt.Printf(
		"Getting block size of %v cipher with key length %v...\n",
		getCipherTypeName(cipherType),
		keyLen,
	)
	blockSize, err := getBlockSize(keyLen, cipherType)
	if err != nil {
		fmt.Println(err)
		fmt.Println("========")
		return
	}
	fmt.Println("Block size:", blockSize)
	fmt.Println("========")
}

func getCipherTypeName(cipherType int) string {
	switch cipherType {
	case typeAES:
		return "AES"
	case typeDES:
		return "DES"
	}
	return "unknown"
}

func main() {
	test(16, typeAES)
	test(24, typeAES)
	test(32, typeAES)
	test(64, typeAES)

	test(8, typeDES)
	test(16, typeDES)
	test(24, typeDES)
	test(1, -1)
}
