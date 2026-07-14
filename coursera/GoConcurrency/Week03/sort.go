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
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {

	// Scanf wont work if your input has spaces :)
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please type in a sequence of integers")
	sequence, _ := consoleReader.ReadString('\n')
	sequence = strings.TrimSuffix(sequence, "\n")

	// Check that a decimal point was included
	r, _ := regexp.Compile("\\.")
	foundDecimal := r.FindString(sequence)
	if foundDecimal != "" {
		panic("Do not input any floating point numbers")
	}

	ints := ConvertStringToInts(sequence)

	size := int(len(ints) / 4)
	if (len(ints) % 4.0) != 0 {
		size++
	}
	partitions := [][]int{}
	partitions = SliceUp(partitions, ints, size)
	// There will be always 4 partions even if some are empty
	for i := 0; i <= (4 - len(partitions)); i++ {
		partitions = append(partitions, []int{})
	}

	var wg sync.WaitGroup
	for i, partition := range partitions {
		wg.Add(1)
		go SortSlice(partition, i, &wg)
	}
	wg.Wait()

	mergedSlice := []int{}
	for _, partition := range partitions {
		mergedSlice = MergeSortedSlices(mergedSlice, partition)
	}
	fmt.Println("All partitions merged and sorted :", mergedSlice)

}

// SortSlice is a wrapper function around BubbleSort
func SortSlice(ints []int, k int, wg *sync.WaitGroup) {
	fmt.Printf("Partition %d before sorting %v\n", k, ints)
	BubbleSort(ints)
	fmt.Printf("Partition %d after sorting %v\n", k, ints)
	wg.Done()
}

// MergeSortedSlices : interleaves two sorted slices and returns a single sorted slice
func MergeSortedSlices(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

// SliceUp : Cuts slice up into partitioned pieces of length size
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

/*
 BubbleSort : modifies a slice so that the elements are in sorted order
  This example is from a previous assignment.
*/
func BubbleSort(ints []int) {
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints)-i-1; j++ {
			if ints[j] > ints[j+1] {
				Swap(ints, j)
			}
		}
	}
}

// Swap to elements in slice. This example is from a previous assignment.
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
