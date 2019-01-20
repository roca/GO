/*
Write a program to sort an array of integers.
The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// Scanf wont work if your input has spaces :)
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please type in a sequence of up to 10 integers")
	sequence, _ := consoleReader.ReadString('\n')
	sequence = strings.TrimSuffix(sequence, "\n")

	ints := ConvertStringToInts(sequence)
	if len(ints) > 10 {
		log.Println("Only 10 intgers please. Try again!")
		os.Exit(1)
	}

	BubbleSort(ints)

	fmt.Print("BubbleSorted: ")
	for _, v := range ints {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")

}
func SliceUp(slices [][]int, ints []int, size int) [][]int {

	if len(ints) == 0 {
		return slices
	}

	partition := []int{}
	start := 0

	for start = 0; start < size && start < len(ints); start++ {
		partition = append(partition, ints[start])
	}
	fmt.Println(partition)
	slices = append(slices, partition)

	return SliceUp(slices, ints[start:], size)

}

// ConvertStringToInts : converts string of ints to a slice of ints
func ConvertStringToInts(s string) []int {

	ints := []int{}
	s = strings.TrimSpace(s)

	stringsArray := regexp.MustCompile("[\\s+|\\,]+").Split(s, -1)
	for _, stringElement := range stringsArray {
		if stringElement == " " {
			continue
		}
		value, _ := strconv.Atoi(stringElement)
		ints = append(ints, value)
	}

	return ints

}

// BubbleSort : modifies a slice so that the elements are in sorted order
func BubbleSort(ints []int) {

	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints)-i-1; j++ {
			if ints[j] > ints[j+1] {
				Swap(ints, j)
			}
		}
	}

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
