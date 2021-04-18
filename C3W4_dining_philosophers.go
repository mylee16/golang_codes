/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{sync.Mutex}
type Philo struct {
	id int
	leftCS, rightCS *ChopS
	times_eaten int
}
var wait sync.Mutex


func (p Philo) eat() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		wait.Lock()
		if (p.times_eaten) < 3 {
			p.leftCS.Lock()
			p.rightCS.Lock()
			
			wait.Unlock()

			p.times_eaten = p.times_eaten + 1
			fmt.Printf("Philosopher %d starting to eat %dtimes\n", p.id, p.times_eaten)
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("Philosopher %d finishing eating \n", p.id)

			p.leftCS.Unlock()
			p.rightCS.Unlock()

			} else {
				wait.Unlock()
			}
		}
	}


var wg sync.WaitGroup

func main(){
	num_philos := 5

	// initiate chopsticks
	CSticks := make([]*ChopS, num_philos)
	for i := 0; i < num_philos; i++ {
		CSticks[i] = new(ChopS)
	}

	// initiate philosophers
	Philos := make([]*Philo, num_philos)
	for i := 0; i < num_philos; i++ {
		Philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5], 0}
}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Philos[i].eat()
	}
	wg.Wait()
	}