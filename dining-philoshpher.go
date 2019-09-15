package main

import (
	"fmt"
	"math"
	"sync"
)

type Chopstick struct{ sync.Mutex }

type Philospher struct {
	Id        int
	LeftChop  *Chopstick
	RightChop *Chopstick
}

func (p Philospher) eat(host chan int, wg *sync.WaitGroup) {

	fmt.Printf("Allowed\n")
	host <- 1
	p.LeftChop.Lock()
	p.RightChop.Lock()

	fmt.Printf("%d Philospher started eating\n", p.Id)

	p.RightChop.Unlock()
	p.LeftChop.Unlock()
	fmt.Printf("%d Philospher completed eating\n", p.Id)
	fmt.Printf("Removed\n")
	<-host
	wg.Done()
}

func main() {
	host := make(chan int, 2)
	var wg sync.WaitGroup
	fmt.Printf("Starting Dinning Philospher problem\n")

	chopsticks := make([]*Chopstick, 5)
	philospher := make([]Philospher, 5)

	for i := range chopsticks {
		chopsticks[i] = new(Chopstick)
	}

	for i := range philospher {
		philospher[i] = Philospher{i, chopsticks[int(math.Min(float64(i), float64((i+1)%5)))], chopsticks[int(math.Max(float64((i+1)%5), float64(i)))]}
	}

	for i := 0; i < 5; i++ {
		for times := 0; times < 3; times++ {
			wg.Add(1)
			go philospher[i].eat(host, &wg)
		}
	}

	wg.Wait()
}
