package main

import (
	"fmt"
	"sync"
)

func main() {
	//buffered channel used as a job queue
	jobs := make(chan int, 5)

	//WaitGroup keeps the main goroutine alive until the worker finishes
	var wg sync.WaitGroup

	wg.Add(1)

	//start a worker goroutine
	go func(wg *sync.WaitGroup) {
		//notify WaitGroup when worker exits
		defer wg.Done()

		//continuously receive jobs until channel closes
		for {
			//receive a job and check whether the channel is still open
			r, ok := <-jobs
			if ok {
				fmt.Println("Got this message", r)
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
	}(&wg)

	//send jobs into the queue
	for i := 1; i <= 3; i++ {
		jobs <- 1
		fmt.Println("Sending", i)
	}

	//signal that no more jobs will be sent
	close(jobs)

	//wait for worker to finish processing
	wg.Wait()
}
