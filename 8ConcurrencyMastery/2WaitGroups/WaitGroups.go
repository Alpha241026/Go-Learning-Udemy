package main

import (
	"fmt"
	"sync"
	"time"
)

// function executed by each goroutine
func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) { //passing by pointer so all goroutines update the same shared counter

	defer wg.Done() //decrease the shared counter when this goroutine finishes

	time.Sleep(delay) //simulating some work

	fmt.Println("sayHello", message)

}

func main() {
	var wg sync.WaitGroup //shared WaitGroup used to synchronize all goroutines

	totalJobs := 5 //no. of concurrent jobs to execute

	for i := 0; i < totalJobs; i++ {
		wg.Add(1)                                               //increment the WaitGroup counter before launching the goroutine
		go sayHello(fmt.Sprintf("JOB %d", i), time.Second, &wg) //start a concurrent job / new goroutine
	}

	fmt.Println("Hello from Main() Goroutine") //every Go program starts with one goroutine: the main goroutine

	wg.Wait() //wait until counter reaches zero
}
