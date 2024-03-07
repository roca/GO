package examples

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Examples_3_16() {
	if len(os.Args[1:]) == 0 {
		fmt.Fprintf(os.Stderr, "%s: %s", time.Now(), "No argument supplied.\n")
		return
	}

	total := 1
	for _, in := range os.Args[1:] {
		i, err := strconv.Atoi(in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s, %s (%s)", time.Now(), "wrong argument supplied", err, in)
			return
		}

		total = total * i
	}

	fmt.Println(total)
}
