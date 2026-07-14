package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// numberOfPartitions also influence the number of goroutines started
const numberOfPartitions = 4

// shallowSort is kind of Merge sort (https://en.wikipedia.org/wiki/Merge_sort) ;)
// It does not split the integers into sublist with single element (as merge sort does) but into several groups.
// Number of groups is defined in numberOfPartitions constant.
// Because it doesn't go so deep with split, thus the name Shallow sort :)
func shallowSort(ints []int) {
	// The partitions is the slice of slices using a single underlying array.
	// Even when the same array is modified from several goroutines
	// there is no risk of a race condition because of goroutines
	// are modifying non-overlapping parts of an array (slices).
	// If the slices would overlap then there will be a race.
	// Try to run the code with Data Race Detector (https://golang.org/doc/articles/race_detector.html)
	// to see that I'm not kidding ;)
	partitions := split(ints, numberOfPartitions)

	var wg sync.WaitGroup
	wg.Add(numberOfPartitions)
	for i := 0; i < numberOfPartitions; i++ {
		go func(gn int, ints []int) {
			defer wg.Done()
			fmt.Printf("goroutine(%d) going to sort: %v\n", gn, ints)
			sort.Ints(ints)
		}(i, partitions[i])
	}
	wg.Wait()
	fmt.Printf("Before merge: %v\n", partitions)

	// Because merged is still sharing a single underlying array with ints and partitions
	// as soon as the merged is sorted the ints and partitions are sorted as well.
	merged := partitions[0]
	for i := 1; i < numberOfPartitions; i++ {
		// It does not make a sense to merge like this but if the slice would be sorted in one step
		// somebody may ask why it was split few lines above ;)
		merged = append(merged, partitions[i]...)
		sort.Ints(merged)
	}
}

func split(ints []int, parts int) [][]int {
	d := len(ints) / parts
	splitted := make([][]int, parts)
	for i := 0; i < parts; i++ {
		first := i * d
		last := (i * d) + d
		if i == parts-1 {
			// Last part, if there is anything remaining add it as well
			last += len(ints) % parts
		}
		splitted[i] = ints[first:last]
	}
	return splitted
}

func input(label string) string {
	fmt.Printf("%s: ", label)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func main() {
	in := input("Enter integers (separated by space) to sort (non-integers will be skipped)")
	vals := strings.Fields(in)
	ints := make([]int, 0, len(vals))
	for _, v := range vals {
		i, err := strconv.Atoi(v)
		if err != nil {
			// Skip the non-integer input
			continue
		}
		ints = append(ints, i)
	}
	if len(ints) < 1 {
		log.Fatalln("No integer entered!")
	}

	fmt.Println("Going to sort: ", ints)
	shallowSort(ints)
	fmt.Println("Sorted: ", ints)
	fmt.Println("Thanks.")
}
