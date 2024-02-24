package main

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
)

// Generate two large prime numbers
func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p,_ := getBigPrime(keysize)
	q,_ := getBigPrime(keysize)

	return p, q
}

// Calculate n = p * q
func getN(p, q *big.Int) *big.Int {
	r := new(big.Int).Set(p)
	s := new(big.Int).Set(q)

	return r.Mul(r, s)
}

// don't touch below this line

func firstNDigits(n big.Int, numDigits int) string {
	if len(n.String()) < numDigits {
		return fmt.Sprintf("%v", n.String())
	}
	return fmt.Sprintf("%v...", n.String()[:numDigits])
}

func test(keySize int) {
	p, q := generatePrivateNums(keySize)
	fmt.Printf("Generated p: %v it has %v digits\n", firstNDigits(*p, 10), len(p.String()))
	fmt.Printf("Generated q: %v it has %v digits\n", firstNDigits(*q, 10), len(q.String()))

	n := getN(p, q)
	fmt.Printf("Generated n: %v it has %v digits\n", firstNDigits(*n, 10), len(n.String()))

	fmt.Println("========")
}

var randReader = mrand.New(mrand.NewSource(0))

func getBigPrime(bits int) (*big.Int, error) {
	if bits < 2 {
		return nil, errors.New("prime size must be at least 2-bit")
	}
	b := uint(bits % 8)
	if b == 0 {
		b = 8
	}
	bytes := make([]byte, (bits+7)/8)
	p := new(big.Int)
	for {
		if _, err := io.ReadFull(randReader, bytes); err != nil {
			return nil, err
		}
		bytes[0] &= uint8(int(1<<b) - 1)
		if b >= 2 {
			bytes[0] |= 3 << (b - 2)
		} else {
			bytes[0] |= 1
			if len(bytes) > 1 {
				bytes[1] |= 0x80
			}
		}
		bytes[len(bytes)-1] |= 1
		p.SetBytes(bytes)
		if p.ProbablyPrime(20) {
			return p, nil
		}
	}
}

func main() {
	test(5)
	test(512)
	test(1024)
	test(2048)
}
