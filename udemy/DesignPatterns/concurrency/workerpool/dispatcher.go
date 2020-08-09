package main

import "time"

type IDispatcher interface {
	LaunchWorker(w IWorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(w IWorkerLauncher) {
	w.LaunchWorker(d.inCh)
}
func (d *dispatcher) Stop() {
	close(d.inCh)
}
func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func NewDispatcher(b int) IDispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}
