package main

import "testing"

func TestBubbleSort(t *testing.T) {

}

func TestSwap(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}

	Swap(ints, 0)
	if ints[0] != 2 && ints[1] != 1 {
		t.Error("Swap positons of 0 with 1 was not correct")
	}

	Swap(ints, 4) // Swap last two
	if ints[3] != 5 && ints[4] != 4 {
		t.Error("Swap(4) positons of 4 with 5 was not correct")
	}
	Swap(ints, 3) // Swap back last two another way
	if ints[3] != 4 && ints[4] != 5 {
		t.Error("Swap positons of 0 with 1 was not correct")
	}

	ints2 := []int{}
	err := Swap(ints2, 8) // Shouldn't do anything
	if err.Error() != "no vaules to swap! slice size <= 1" {
		t.Error(err)
	}
}
