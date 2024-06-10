package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker",id, "started job", j, "...")
		time.Sleep(time.Second)
		fmt.Println("Worker",id, "finished job", j, "...")
		results <- j * 2
	}
}

func main() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for id := 1; id <= 3; id++ {
		go worker(id, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
