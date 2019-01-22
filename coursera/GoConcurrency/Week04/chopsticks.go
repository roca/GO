package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

var on sync.Once
var wg sync.WaitGroup

func setup() {
	//wg.Add(5)
	fmt.Println("Init")
}

func (p Philo) eat() {

	on.Do(setup)

	for {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}

	//wg.Done()
}

func main() {
	Csticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		Csticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{Csticks[i], Csticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}

	// wg.Wait()
}
