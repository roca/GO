package main

import (
	"fmt"
	"sync"
)

// 1. There are 5 philosophers and 5 chopsticks.
// 2. There is one host channel which holds two tokens.

// Before eating:
// Philosopher gets the token from the host.
// The host is a buffered channel with 2 tokens in it.
// Once the philosopher has a token it will attempt to take chopsticks.
// Here even when two neighboring philosophers are holding tokens,
// at least one of then will be able to pick two both chopsticks
// and continue. Thus at least one philosopher can eat at any given time.

// After Eating: philosopher releases chopsticks first.
// So, anyone with a token but waiting on chopstick can continue.
// After that philosopher will give the token back to the host.
// Thus, anyone waiting on token can now continue and pick chopsticks.

// The host has only two tokens so third person waiting for a token will be waiting until someone releases the token.

var wg sync.WaitGroup

const token1 = 1
const token2 = 2

var host = make(chan int, 2)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	count     int
	chopleft  *ChopStick
	chopright *ChopStick
}

func (phil Philosopher) eat() {
	for i := 0; i < 3; i++ {
		token := <-host
		phil.chopleft.Lock()
		phil.chopright.Lock()

		// phil.count++
		fmt.Println("starting to eat", phil.count)
		fmt.Println("finishing to eat", phil.count)

		phil.chopright.Unlock()
		phil.chopleft.Unlock()
		host <- token
	}
	wg.Done()
}

func main() {
	var chopsticks []*ChopStick = make([]*ChopStick, 5)
	var philosophers []*Philosopher = make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		chopsticks[i] = &ChopStick{}
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			count:     i + 1,
			chopleft:  chopsticks[(i+1)%5],
			chopright: chopsticks[i],
		}
	}
	host <- token1
	host <- token1

	wg.Add(5)
	for _, p := range philosophers {
		go p.eat()
	}

	wg.Wait()
}
