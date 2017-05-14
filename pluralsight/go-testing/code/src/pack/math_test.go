package pack

import "testing"
import "time"

func TestCanAddNumbers(t *testing.T) {
	result := Add(1, 2)
	time.Sleep(3 * time.Second)
	if result != 3 {
		t.Log("Failed to add one and two")
		t.Fail()
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
