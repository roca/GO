/*

Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.
The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1.
I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//float64 arguments, acceleration a, initial velocity vo, and initial displacement so.

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {

	return func(t float64) float64 {
		return (.5 * a * math.Pow(t, 2.0)) + (vo * t) + so
	}
}

// ConvertStringToFloats : converts string of numbers to a slice of floats
func ConvertStringToFloats(s string) []float64 {

	floats := []float64{}
	s = strings.TrimSpace(s)

	stringsArray := regexp.MustCompile("[\\s+|\\,]+").Split(s, -1)
	for _, stringElement := range stringsArray {
		if stringElement == " " {
			continue
		}
		value, _ := strconv.ParseFloat(stringElement, 64)
		floats = append(floats, value)
	}

	return floats

}

func main() {

	// Scanf wont work if your input has spaces :)
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter values for acceleration, initial velocity, and initial displacement on one line. (etc. 10, 2, 1")
	params, _ := consoleReader.ReadString('\n')
	params = strings.TrimSuffix(params, "\n")

	floats := ConvertStringToFloats(params)
	if len(floats) > 3 {
		log.Println("Only 3 values are needed. Try again!")
		os.Exit(1)
	}

	var time float64

	fmt.Println("Now enter a value for time.")
	fmt.Scan(&time)

	fmt.Println(floats)
	fmt.Println(time)

	fn := GenDisplaceFn(floats[0], floats[1], floats[2])

	fmt.Println("displacements = ½ a t2 + vot + so = ", fn(time))

}
