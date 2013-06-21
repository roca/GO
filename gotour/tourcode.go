//package main

//import (
//	"fmt"
//	"math"
//	"time"
//)

//type MyError struct {
//	When time.Time
//	What string
//}

//func (e *MyError) Error() string {
//	return fmt.Sprintf("at %v, %s",
//		e.When, e.What)
//}

//func run() error {
//	return &MyError{
//		time.Now(),
//		"it didn't work",
//	}
//}

//func Cbrt(x complex128) (complex128, error) {
//	//if real(x) < 0 {
//	//return 0, run()
//	//}
//	isGoodEnough := func(guess complex128) bool {
//		z := (guess * guess * guess) - x
//		e := math.Sqrt((real(z) * real(z)) + (imag(z) * imag(z)))
//		return e < 0.0001
//	}
//	improve := func(guess complex128) complex128 {
//		return guess - (((guess * guess * guess) - x) / (3 * guess * guess))
//	}
//	cbrtItr := func(guess complex128) complex128 { return complex(0, 0) }
//	cbrtItr = func(guess complex128) complex128 {

//		if isGoodEnough(guess) {
//			return guess
//		} else {
//			return cbrtItr(improve(guess))
//		}
//	}
//	return cbrtItr(complex(1.0, 1.0)), nil

//}

//func main() {
//	fmt.Println(Cbrt(-2))
//}

package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
