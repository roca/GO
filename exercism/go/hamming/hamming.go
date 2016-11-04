package hamming

import (
	"errors"
	"strings"
)

const testVersion = 5

// Distance performs hamming calulations
func Distance(a, b string) (int, error) {
	// your code here
	aa := strings.Split(a, "")
	bb := strings.Split(b, "")

	if len(aa) != len(bb) {
		return 0, errors.New("Sequences have different lengths")
	}

	score := 0

	for i, aBase := range aa {
		if i < len(bb) && aBase != bb[i] {
			score++
		}
	}

	return score, nil

}
