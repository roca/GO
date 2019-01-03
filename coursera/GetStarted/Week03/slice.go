package main

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	sliceInputs := make([]int, 0, 3)
	var newElement string

	for {
		fmt.Print("\nPlease enter an integer to be added to the slice. Type an 'X' when your done :  ")
		fmt.Scan(&newElement)
		i, err := GetIntFromString(newElement)
		if err != nil {
			log.Println(err)
			break
		}
		sliceInputs = append(sliceInputs, i)
		sort.Ints(sliceInputs)
		fmt.Println(sliceInputs)
	}

}

// GetIntFromString ...
func GetIntFromString(v string) (int, error) {

	//	Trim away Leading and Trailing spaces
	v = strings.TrimSpace(v)

	// Check that a decimal point was not included
	regex, err := regexp.Compile("(?i)X")
	if regex.MatchString(v) {
		return 0, errors.New("Done")
	}

	// Input should parse to a Intege
	i, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}
