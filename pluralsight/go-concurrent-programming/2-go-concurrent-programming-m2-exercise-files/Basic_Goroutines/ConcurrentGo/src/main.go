package main

import (
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(2)
	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
			time.Sleep(godur)

		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)

}

/*
	2.1.1 - two funcs called in sequence
	func(){
		println("Hello")
	}()

	func() {
		println("Go")
	}()
*/

/*
	2.1.2 - main thread exits before go routines have chance to run
	go func() {
		println("Hello")
	}()

	go func() {
		println("Go")
	}()
*/

/*
	2.1.3 - add sleep to main thread to allow go routines to execute
	go func() {
		println("Hello")
	}()

	go func() {
		println("Go")
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
*/

/*
	2.1.4 - goroutines will swap execution contexts
	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i:=0; i< 100; i++ {
			println("Hello")
			time.Sleep(godur)

		}
	}()

	go func() {
		for i:=0; i< 100; i++ {
			println("Go")
			time.Sleep(godur)
		}
	}()


	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
*/

/*
	2.1.5 - with multiple processor threads, application becomes nondeterminant
	runtime.GOMAXPROCS(2)
	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i:=0; i< 100; i++ {
			println("Hello")
			time.Sleep(godur)

		}
	}()

	go func() {
		for i:=0; i< 100; i++ {
			println("Go")
			time.Sleep(godur)
		}
	}()


	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
*/
