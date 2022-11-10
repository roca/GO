package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	op := flag.String("op", "sum", "Operation to be executed: sum")
	column := flag.Int("col", 1, "CSV column on which to execute the operation")
	flag.Parse()
	if err := run(flag.Args(), *op, *column, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filenames []string, op string, column int, out io.Writer) error {
	var opFunc statsFunc

	if len(filenames) == 0 {
		return ErrNoFiles
	}

	if column < 1 {
		return fmt.Errorf("%w: %d", ErrInvalidColumn, column)
	}

	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}

	consolidated := make([]float64, 0)

	for _, fname := range filenames {
		f, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("Cannot open file %s: %w", fname, err)
		}

		data, err := csv2float(f, column)
		if err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}

		consolidated = append(consolidated, data...)

	}

	_, err := fmt.Fprintln(out, opFunc(consolidated))
	return err

}
