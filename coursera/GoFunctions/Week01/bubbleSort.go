/*
Write a Bubble Sort program in Go.
The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
*/
package main

import "errors"

func main() {

}

func BubbleSort(ints []int) {

}

// Swap to elements in slice
func Swap(ints []int, position int) error {
	if len(ints) <= 1 {
		return errors.New("no vaules to swap! slice size <= 1")
	}
	if position > len(ints)-1 {
		return errors.New("position index out of range")
	}

	if position != len(ints)-1 {
		ints[position], ints[position+1] = ints[position+1], ints[position]
	} else { // Swaps the last two element if position is the last element
		ints[position-1], ints[position] = ints[position], ints[position-1]
	}

	return nil
}
