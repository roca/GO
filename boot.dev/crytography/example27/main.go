package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
)

// Calculate ϕ(n) = (p-1)(q-1)
func getTot(p, q *big.Int) *big.Int {
	tot := new(big.Int)
	tot.Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	return tot
}

// Choose a public exponent
func getE(tot *big.Int) *big.Int {
	totMinusTwo := new(big.Int)
	totMinusTwo.Sub(tot, big.NewInt(2))

	e, _ := rand.Int(randReader, totMinusTwo)
	e.Add(e, big.NewInt(2))
	for gcd(e, tot).Cmp(big.NewInt(1)) != 0 {
		e, _ = rand.Int(randReader, totMinusTwo)
		e.Add(e, big.NewInt(2))
	}
	return e
}

// don't touch below this line

func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p, _ := getBigPrime(keysize)
	q, _ := getBigPrime(keysize)
	return p, q
}

func getN(p, q *big.Int) *big.Int {
	n := new(big.Int)
	n.Mul(p, q)
	return n
}

func gcd(x, y *big.Int) *big.Int {
	xCopy := new(big.Int).Set(x)
	yCopy := new(big.Int).Set(y)
	for yCopy.Cmp(big.NewInt(0)) != 0 {
		xCopy, yCopy = yCopy, xCopy.Mod(xCopy, yCopy)
	}
	return xCopy
}

func firstNDigits(n big.Int, numDigits int) string {
	if len(n.String()) < numDigits {
		return fmt.Sprintf("%v", n.String())
	}
	return fmt.Sprintf("%v...", n.String()[:numDigits])
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

func test(keySize int) {
	p, q := generatePrivateNums(keySize)
	fmt.Printf("Generated p: %v it has %v digits\n", firstNDigits(*p, 10), len(p.String()))
	fmt.Printf("Generated q: %v it has %v digits\n", firstNDigits(*q, 10), len(q.String()))

	n := getN(p, q)
	fmt.Printf("Generated n: %v it has %v digits\n", firstNDigits(*n, 10), len(n.String()))

	tot := getTot(p, q)
	fmt.Printf("Generated tot: %v it has %v digits\n", firstNDigits(*tot, 10), len(tot.String()))

	e := getE(tot)
	fmt.Printf("Generated e: %v it has %v digits\n", firstNDigits(*e, 10), len(e.String()))
	fmt.Println("========")
}

func main() {
	test(512)
	test(1024)
}