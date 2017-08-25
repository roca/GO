package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	flag.String("h", "help", "Method of Least Squares")
	var nPtr = flag.Int("n", 1, "order of eqution y = AX^n + BX^(n-1) + .... CX^0")
	flag.Parse()

	nOrder := *nPtr

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

	fmt.Println(m)

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
