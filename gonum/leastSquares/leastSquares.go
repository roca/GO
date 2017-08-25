package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/gonum/matrix/mat64"
)

func main() {

	flag.String("h", "help", "Method of Least Squares")
	var nPtr = flag.Int("n", 1, "order of eqution y = AX^n + BX^(n-1) + .... CX^0")
	flag.Parse()

	nOrder := *nPtr + 1

	if len(os.Args) == 1 {
		fmt.Printf(usage(filepath.Base(os.Args[0])))
		os.Exit(1)
	}
	dataFilePath := os.Args[len(os.Args)-1]
	_, err := os.Stat(dataFilePath)
	if err != nil {
		fmt.Println(MyError{fmt.Sprintf("NonExisting file path : %s", dataFilePath)})
		fmt.Printf(usage(filepath.Base(os.Args[0])))
		os.Exit(1)
	}
	fmt.Printf("Data points in %s will be fit to %d order polynomial\n", dataFilePath, *nPtr)

	xx := make([][]float64, nOrder)
	xy := make([]float64, nOrder)
	for i := range xx {
		xx[i] = make([]float64, nOrder)
	}

	f, err := os.Open(dataFilePath)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {

		re := regexp.MustCompile("  +")
		data := re.ReplaceAllString(s, "")
		point := strings.Split(data, ",")
		pointX, errX := strconv.ParseFloat(point[0], 64)
		if errX != nil {
			fmt.Println("Can't read %s", point[0])
		}
		pointY, errY := strconv.ParseFloat(point[1], 64)
		if errY != nil {
			fmt.Println("Can't read %s", point[1])
		}
		// fmt.Printf("%g %g\n", pointX, pointY)

		for i, v1 := range xx {
			xy[i] = xy[i] + (pointY * math.Pow(pointX, float64(i)))
			for j, _ := range v1 {
				xx[i][j] = xx[i][j] + math.Pow(pointX, (float64(i)+float64(j)))
			}
		}

		s, e = Readln(r)
	}

	// for i, v1 := range xx {
	// 	for j, v2 := range v1 {
	// 		fmt.Printf("(%d %d) %g ", i, j, v2)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// for i, v1 := range xy {
	// 	fmt.Printf("(%d) %g ", i, v1)
	// 	fmt.Printf("\n")
	// }

	mXY := mat64.NewDense(nOrder, 1, xy)
	solution := mat64.NewDense(nOrder, 1, nil)

	// print all mXY elements
	fmt.Printf("mXY :\n%v\n\n", mat64.Formatted(mXY, mat64.Prefix(""), mat64.Excerpt(0)))

	mXX := mat64.NewDense(nOrder, nOrder, nil)
	for i, _ := range xy {
		mXX.SetRow(i, xx[i])
	}

	// print all mXX elements
	fmt.Printf("mXX :\n%v\n\n", mat64.Formatted(mXX, mat64.Prefix(""), mat64.Excerpt(0)))

	solution.Solve(mXX, mXY)

	fmt.Printf("solution :\n%v\n\n", mat64.Formatted(solution, mat64.Prefix(""), mat64.Excerpt(0)))

}

func usage(arg string) string {
	return fmt.Sprintf("usage: %s -n=[5|4|3...1] <path_of_data_points_file>\n", arg)
}

// MyError is an error implementation that includes a time and message.
type MyError struct {
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v", e.What)
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
