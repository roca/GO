package main

import (
	"fmt"
	//"os"
	"time"
)

func main() {

	boring("boring!")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
