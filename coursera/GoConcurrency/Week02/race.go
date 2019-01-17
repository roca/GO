package main

import (
	"fmt"
	"time"
)

/*
Outcome of x depends on non-deterministic ordering. If you run this program several times
in succession you will see that the value of x will change depending on how the go schedular
decides to order there execution. A race condition occurs in this case because both
go routines are accessing the same piece of data, x in this case. Depending on how the schedular decides to
execute these two routines the final outcome of x may be 0 or it maybe 1.
*/

func main() {
	x := 0

	go func() {
		x++
	}()

	go func() {
		fmt.Println(x)
	}()

	time.Sleep(100 * time.Millisecond)

}
