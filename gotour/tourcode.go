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

//package main

//import (
//	"fmt"
//	"net/http"
//)

//type Hello struct{}

//type String string

//type Struct struct {
//	Greeting string
//	Punct    string
//	Who      string
//}

//func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprint(w, s)
//}

//func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	message := fmt.Sprintf("%s%s%s", s.Greeting, s.Punct, s.Who)
//	fmt.Fprint(w, message)
//}

//func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprint(w, "Hello!")
//}

//func main() {
//	var h Hello
//	http.Handle("/string", String("I'm a frayed knot."))
//	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
//	http.Handle("/", h)
//	http.ListenAndServe("localhost:4000", nil)
//}

// Excericize #58
//

package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (I *Image) ColorModel() color.Model { return color.RGBAModel }
func (I *Image) Bounds() image.Rectangle { return image.Rect(0, 0, 40, 40) }
func (I *Image) At(x, y int) color.Color { return color.RGBA{uint8(x), uint8(y), 255, 255} }

func main() {
	m := &Image{}
	pic.ShowImage(m)
}
