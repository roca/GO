package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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
	fmt.Printf("Data points in %s will be fit to %d order polynomial\n", dataFilePath, nOrder)

	m := make([][]float64, nOrder)
	for i := range m {
		m[i] = make([]float64, nOrder)
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
		fmt.Printf("%g %g\n", pointX, pointY)

		for i, v1 := range m {
			for j, _ := range v1 {
				m[i][j] = m[i][j] + (float64(i) + float64(j))
			}
		}

		s, e = Readln(r)
	}

	for i, v1 := range m {
		for j, v2 := range v1 {
			fmt.Printf("(%d %d) %g ", i, j, v2)
		}
		fmt.Printf("\n")
	}

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
