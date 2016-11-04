package hamming

import "errors"

const testVersion = 5

// Distance performs hamming calulations
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("The given sequences have different lengths")
	}

	score := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			score++
		}
	}

	return score, nil

}
