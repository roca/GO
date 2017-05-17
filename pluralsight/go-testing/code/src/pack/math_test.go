package pack

import "testing"
import "time"

func TestCanAddNumbers(t *testing.T) {
	if testing.Short() {
		t.Skip("Skiping long tests")
	}
	result := Add(1, 2)
	time.Sleep(3 * time.Second)
	if result != 3 {
		t.Fatal("Failed to add one and two")
	}

	result = Add(1, 2, 3, 4)
	if result != 10 {
		t.Error("Failed to add more then two numbers")
	}
}

func TestCanSubtractNumber(t *testing.T) {
	result := Subtract(10, 5)

	if result != 5 {
		t.Error("Failed to substract two numbers properly")
	}
}

func TestCanMultiplyNumbers(t *testing.T) {
	if testing.Verbose() {
		t.Skip("Not implemented yet")
	}
}
