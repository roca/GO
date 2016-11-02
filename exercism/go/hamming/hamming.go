package hamming

import (
	"fmt"
	"strings"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
	// your code here
	aa := strings.Split(a, "")
	bb := strings.Split(b, "")

	score := 0

	for i, a_base := range aa {
		if a_base != bb[i] {
			score++
		}

	}

	//score = 0

	fmt.Printf("score %d\n", score)
	return score, nil

}
