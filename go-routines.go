package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	/**
	Logging for the new worker
	*/
	fmt.Printf("Worker %d starting\n", id)

	/**
	Faking the heavy process
	*/
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)

	/**
	messaging that this worker is completed
	*/
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		/**
		Adding a new worker in the waiting list
		*/
		wg.Add(1)

		/**
		go routine to start a new worker
		*/
		go worker(i, &wg)
	}

	wg.Wait()
}

/**
several output of this experiment

output 1:
	Worker 2 starting
	Worker 1 starting
	Worker 2 done
	Worker 1 done

output 2:
	Worker 1 starting
	Worker 2 starting
	Worker 1 done
	Worker 2 done

output 3:
	Worker 1 starting
	Worker 2 starting
	Worker 1 done
	Worker 2 done
*/
