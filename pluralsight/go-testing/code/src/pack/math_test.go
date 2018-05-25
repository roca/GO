package pack

import (
	"os"
	"testing"
)

var addTable = []struct {
	in  []int
	out int
}{
	{[]int{1, 2}, 3},
	{[]int{1, 2, 3, 4}, 10},
}

func TestMain(m *testing.M) {
	println("Tests are about to run")
	result := m.Run()
	println("Tests done executing")

	os.Exit(result)
}

func TestCanAddNumbers(t *testing.T) {
	t.Parallel()
	for _, entry := range addTable {
		result := Add(entry.in...)
		if result != entry.out {
			t.Error("Failed to add numbers as expected")
		}
	}
}

func TestCanSubtractNumber(t *testing.T) {
	t.Parallel()
	result := Subtract(10, 5)
	if result != 5 {
		t.Error("Failed to substract two numbers properly")
	}
}

func TestCanMultiplyNumbers(t *testing.T) {
	//if testing.Verbose() {
	t.Skip("Not implemented yet")
	//}
}
