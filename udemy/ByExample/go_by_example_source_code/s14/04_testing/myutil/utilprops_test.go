// File name: ...\s14\04_testing\myutil\utilprops_test.go

package myutil

import (
	"fmt"
	"testing"
)

type rectTestData struct {
	l     uint
	w     uint
	area  uint
	perim uint
}

// change any value in the following test set to experience a fail example.
var rectTestsValues = []rectTestData{
	{2, 3, 6, 10}, {5, 10, 50, 30}, {4, 2, 8, 12}}

func TestRect(t *testing.T) {
	for _, el := range rectTestsValues { //el for element
		tmpArea := Area(el.l, el.w)
		tmpPerim := Perim(el.l, el.w)

		if tmpArea != el.area {
			t.Errorf("[Test Name: %s * Input: %d, %d * Expected: %d * Got: %d]\n",
				t.Name(), el.l, el.w, el.area, tmpArea)
		}

		if tmpPerim != el.perim {
			t.Errorf("[Test Name: %s * Input: %d, %d * Expected: %d * Got: %d]\n",
				t.Name(), el.l, el.w, el.perim, tmpPerim)
		}
	}

	fmt.Printf("** TestRect - ALL PASSED (Number of test cases: %d) **\n", len(rectTestsValues))
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
type avgTestData struct {
	val []float32
	avg float32
}

var avgTestsValues = []avgTestData{
	{[]float32{4, 6}, 5.00},
	{[]float32{3, -3}, 0},
	{[]float32{3, -3, 5, -5, 4}, 0.8},
	{[]float32{1, 1, 1, 100}, 25.75},
}

func TestAverage(t *testing.T) {
	for _, el := range avgTestsValues {
		tmpAvg := Avg(el.val)
		if tmpAvg != el.avg {

			t.Errorf("[Test Name: %s * Input: %.2f * Expected: %.2f * Got: %.2f]\n",
				t.Name(), el.val, el.avg, tmpAvg)
		}
	}
	fmt.Printf("** TestAverage - ALL PASSED (Number of test cases: %d) **\n", len(avgTestsValues))
}
