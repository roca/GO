package main

import (
	"fmt"
	"math"
)

// This example is only to demonstrate how hashtables work.
// The logic to make HashKeys is not optimized.
// This solution only works for up to 3 characters.
func main() {
	values := []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}

	// 65x100 + 66x10 + 67x1 = 7227

	hashMap := map[int]string{}
	for _, v := range values {

		hashMap[HashK(v)] = v

	}

	fmt.Println(hashMap)
}

func HashK(s string) int {
	return HashKey(len(s)-1, 0, s)
}

func HashKey(i int, key int, chars string) int {
	if i == 0 {
		return key + int(chars[0])*int(math.Pow10(i))
	}
	return HashKey(i-1, key+int(chars[0])*int(math.Pow10(i)), chars[1:])
}
