package main

import (
	"testing"
)

func TestSliceUp(t *test.T) {
	ints := []int{1, 3, 2, 4, 5, 9, 10, 6, 8, 7, 0}
	slices := SliceUp(ints)

}

func TestBubbleSort(t *testing.T) {

	ints := []int{1, 3, 2, 4, 5, 9, 10, 6, 8, 7, 0}
	BubbleSort(ints)
	for i, v := range ints {
		if v != i {
			t.Errorf("%d != %d\n", v, i)
		}
	}

}

func TestConvertStringToInts(t *testing.T) {
	ints := ConvertStringToInts("0 1 2 3 4 ")
	if len(ints) != 5 {
		t.Errorf("ints is oncrrect length %d insted of 5", len(ints))
	}

	for i, v := range ints {
		if v != i {
			t.Errorf("%d != %d\n", v, i)
		}
	}

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

	ints3 := []int{1, 2, 3}
	err = Swap(ints3, 8) // index out of range
	if err.Error() != "position index out of range" {
		t.Error(err)
	}
}
