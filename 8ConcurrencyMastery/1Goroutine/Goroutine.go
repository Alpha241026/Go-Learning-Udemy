package main

import (
	"fmt"
	"time"
)

// function executed by each goroutine
func sayHello(message string, delay time.Duration) {
	time.Sleep(delay) //simulating some work
	fmt.Println("sayHello", message)
}

func main() {
	fmt.Println("Hello from Main() Goroutine") //every Go program starts with one goroutine: the main goroutine

	//independent goroutines
	go sayHello("Hello World 1", time.Second)
	go sayHello("Hello World 2", time.Second)
	go sayHello("Hello World from 2 seconds", 2*time.Second)
	go sayHello("Hello World from 3 seconds", 3*time.Second)

	//main continues immediately without waiting
	fmt.Println("Last message from Main() Goroutine")

	//keeping program alive temporarily...since all other goroutines stop when main exits
	time.Sleep(2 * time.Second)
}
