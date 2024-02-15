package main

import "fmt"

func crypt(plaintext, key []byte) []byte {
	var xor []byte
	for i := range plaintext {
		xor = append(xor, plaintext[i]^key[i])
	}
	return xor
}

func toString(bytes []byte) string {
	s := ""
	for _, v := range bytes {
		s += fmt.Sprintf("%d", v)
	}
	return s
}

// don't touch below this line

func encrypt(plaintext, key []byte) []byte {
	return crypt(plaintext, key)
}

func decrypt(ciphertext, key []byte) []byte {
	return crypt(ciphertext, key)
}

func test(plaintext, key []byte) {
	ciphertext := encrypt(plaintext, key)
	fmt.Printf("Encrypting '%s' using key '%s'\n", string(plaintext), string(key))
	fmt.Printf("Encrypted ciphertext bytes: %v\n", ciphertext)
	fmt.Printf("Encrypted ciphertext string: %s\n", toString(ciphertext))
	decrypted := decrypt(ciphertext, key)
	fmt.Printf("Decrypted message: %v\n", string(decrypted))
	fmt.Println("========")
}

func main() {
	test([]byte("Shazam"), []byte("Sk7p13"))
	test([]byte("I'm lovin it"), []byte("mysecurepass"))
	test([]byte("Don't tell him I'm in love"), []byte("c5f149783abf22a96e9a7bb999"))
}
