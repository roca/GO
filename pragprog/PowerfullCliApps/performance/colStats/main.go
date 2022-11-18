package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
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
	resCh := make(chan []float64)
	errCh := make(chan error)
	doneCh := make(chan struct{})
	filesCh := make(chan string)

	go func() {
		defer close(filesCh)
		for _, filename := range filenames {
			filesCh <- filename
		}
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for fname := range filesCh {
				f, err := os.Open(fname)
				if err != nil {
					errCh <- fmt.Errorf("Cannot open file %s: %w", fname, err)
					return
				}

				data, err := csv2float(f, column)
				if err != nil {
					errCh <- err
				}

				if err := f.Close(); err != nil {
					errCh <- err
				}

				resCh <- data
			}
		}()
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case data := <-resCh:
			consolidated = append(consolidated, data...)
		case <-doneCh:
			_, err := fmt.Fprintln(out, opFunc(consolidated))
			return err
		}
	}
}
