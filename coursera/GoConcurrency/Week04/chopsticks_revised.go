/*
1 There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2 Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3 The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4 In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5 The host allows no more than 2 philosophers to eat concurrently.
6 Each philosopher is numbered, 1 through 5.
7 When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself,
  where <number> is the number of the philosopher.
8 When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself,
  where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	id          int
	stick       sync.Mutex
	isBeingUsed bool
}

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
	eatCount        int
}

var on sync.Once
var wg sync.WaitGroup

/*
	eat(true): philosophers pick up chopsicks in a random order.
		A deadlock can occur in this case.

	eat(false): philosophers pick up chopsicks in a non-random order.
		The chopsick with the lowest id gets picked up first.
		Seems to avoid deadlocks in this case.
*/
func (p *Philo) eat(randomOrder bool) {

	for i := 0; i < 3; i++ {

		if p.leftCS.id < p.rightCS.id && !randomOrder {
			p.leftCS.stick.Lock()
			p.leftCS.isBeingUsed = true
			p.rightCS.stick.Lock()
			p.rightCS.isBeingUsed = true
		} else if p.leftCS.id >= p.rightCS.id && !randomOrder {
			p.rightCS.stick.Lock()
			p.rightCS.isBeingUsed = true
			p.leftCS.stick.Lock()
			p.leftCS.isBeingUsed = true
		} else {
			p.leftCS.stick.Lock()
			p.leftCS.isBeingUsed = true
			p.rightCS.stick.Lock()
			p.rightCS.isBeingUsed = true
		}

		fmt.Printf("\tphilosopher %d starting to eat with copstick left %d and right %d\n", p.id, p.leftCS.id, p.rightCS.id)
		p.eatCount++

		time.Sleep(1000 * time.Millisecond) // philosophers need time to eat too :);

		if p.leftCS.id < p.rightCS.id && !randomOrder {
			p.leftCS.isBeingUsed = false
			p.leftCS.stick.Unlock()
			p.rightCS.isBeingUsed = false
			p.rightCS.stick.Unlock()
		} else if p.leftCS.id >= p.rightCS.id && !randomOrder {
			p.rightCS.isBeingUsed = false
			p.rightCS.stick.Unlock()
			p.leftCS.isBeingUsed = false
			p.leftCS.stick.Unlock()
		} else {
			p.leftCS.isBeingUsed = false
			p.leftCS.stick.Unlock()
			p.rightCS.isBeingUsed = false
			p.rightCS.stick.Unlock()
		}

		fmt.Printf("\tphilosopher %d finishing eating with copstick left %d and right %d\n", p.id, p.leftCS.id, p.rightCS.id)

	}
	wg.Done()

}

func main() {

	Csticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		stick := new(sync.Mutex)
		Csticks[i] = &ChopS{i, *stick, false}
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, Csticks[i], Csticks[(i+1)%5], 0}
	}

	fmt.Println("Each philosopher eats three times")
	for _, philosopher := range philos {
		HostPhilosopher(philosopher)
	}
	wg.Wait()

	for _, philosopher := range philos {
		fmt.Printf("Philosopher %d eat %d times\n", philosopher.id, philosopher.eatCount)
	}
}

// HostPhilosopher ...
func HostPhilosopher(philosopher *Philo) {
	philosopher.eatCount = 0
	wg.Add(1)
	/*
		 eat(true): philosophers pick up chopsicks in a random order.
			   A deadlock can occur in this case.

		 eat(false): philosophers pick up chopsicks in a non-random order.
				The chopsick with the lowest id gets picked up first.
				Seems to avoid deadlocks in this case.
	*/
	go philosopher.eat(true) // true: philosophers pick up chopsicks in a random order

}
