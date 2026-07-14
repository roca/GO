package interpreter

import "testing"

func TestCalculate(t *testing.T) {
	tempOperation := "3 3 sum 2 sub"
	res, err := Calculate(tempOperation)
	if err != nil {
		t.Error(err)
	}
	if res != 4 {
		t.Errorf("Expected result not found: %d != %d\n", 4, res)
	}
	tempOperation = "3 4 sum 2 sub"
	res, err = Calculate(tempOperation)
	if err != nil {
		t.Error(err)
	}
	if res != 5 {
		t.Errorf("Expected result not found: %d != %d\n", 5, res)
	}
}
