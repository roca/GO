package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
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
	x := make([]float64, 0)
	y := make([]float64, 0)
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

		data := strings.Replace(s, " ", "", -1)
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

		x = append(x, pointX)
		y = append(y, pointY)
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
	fmt.Printf("R2 score: %g \n\n", r2_score(x, y, solution))

}

func usage(arg string) string {
	return fmt.Sprintf("usage: %s -n=[5|4|3...1] <path_of_data_points_file>\n", arg)
}

func r2_score(xs, ys []float64, solution *mat64.Dense) float64 {
	var yMean, sum, ssTotal, predictedY, ssResidual float64
	predictedYs := make([]float64, 0)
	sum = 0
	r, _ := solution.Dims()

	for _, v := range ys {
		sum = sum + v
	}
	yMean = sum / float64(len(ys))

	for _, v := range xs {
		predictedY = 0
		for i := 0; i < r; i++ {
			predictedY = +(solution.At(i, 0) * math.Pow(v, float64(i)))
		}
		predictedYs = append(predictedYs, predictedY)
	}

	ssTotal = 0
	for _, v := range ys {
		ssTotal = ssTotal + math.Pow((v-yMean), 2.0)
	}

	ssResidual = 0
	for i, v := range predictedYs {
		ssResidual = ssResidual + math.Pow((v-ys[i]), 2.0)
	}

	fmt.Printf("Number of points: %d \n\n", len(predictedYs))
	fmt.Printf("ssTotal: %g \n\n", ssTotal)
	fmt.Printf("ssResidual: %g \n\n", ssResidual)

	return 1.0 - (ssResidual / ssTotal)

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
