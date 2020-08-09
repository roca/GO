package main

import (
	"fmt"
	"log"
	"sync"
)

type RequestHandler func(interface{})

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: fmt.Sprintf("(%s: %d) -> Hello",s,id),
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
	return myRequest
}

func main() {
	bufferSize := 100
	var dispatcher IDispatcher = NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		var w IWorkerLauncher = &PrefixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d ->", i),
			suffixS: " World",
			id:      i,
		}
		dispatcher.LaunchWorker(w)
	}

	requests := 10

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest("Msg_id", i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	wg.Wait()
}
