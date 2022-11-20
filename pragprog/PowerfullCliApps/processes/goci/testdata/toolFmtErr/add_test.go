package add

import "testing"

func TestAdd(t *testing.T) {
	a := 2
	b := 3

	exp := 5

	res := add(a, b)

	if res != exp {
		t.Errorf("Expected: %d, Got: %d.", exp, res)
	}
}
