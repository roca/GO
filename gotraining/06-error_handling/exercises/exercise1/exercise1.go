// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/8KETdvYk17

// Create two error variables, one called ErrInvalidValue and the other
// called ErrAmountTooLarge. Provide the static message for each variable.
// Then write a function called checkAmount that accepts a float64 type value
// and returns an error value. Check the value for zero and if it is, return
// the ErrInvalidValue. Check the value for greater than $1,000 and if it is,
// reutrn the ErrAmountTooLarge. Write a main function to call the checkAmount
// function and check the return error value. Display a proper message to the screen.
package main

import (
	"errors"
	"fmt"
)

// ErrInvalidValue indicates the value is invalid.
var ErrInvalidValue = errors.New("Invalid Value")

// ErrAmountTooLarge indicates the value beyond the upper bound.
var ErrAmountTooLarge = errors.New("Amount To Large")

// main is the entry point for the application.
func main() {
	// Call the function and get the error.
	if err := checkAmount(0); err != nil {
		switch err {
		// Check if the error is an ErrInvalidValue.
		case ErrInvalidValue:
			fmt.Println("Value provided is not valid.")
			return

		// Check if the error is an ErrAmountTooLarge.
		case ErrAmountTooLarge:
			fmt.Println("Value provided is too large.")
			return

		// Handle the default error.
		default:
			fmt.Println(err)
			return
		}
	}

	// Display everything is good.
	fmt.Println("Everything checks out.")
}

// checkAmount validates the value passed in.
func checkAmount(f float64) error {
	switch {
	// Is the parameter equal to zero.
	case f == 0:
		return ErrInvalidValue

	// Is the parameter greater than 1000.
	case f > 1000:
		return ErrAmountTooLarge
	}

	return nil
}
