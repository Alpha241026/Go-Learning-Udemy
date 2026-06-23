package main

import "fmt"

func factorial(n int) int { //recursive function to return factorial
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func intSeq() func() int { //returns inner function that increments & returns an integer each time its called
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(factorial(5))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger := func(msg string) { //creating simple inline functtion assigned to a variable
		fmt.Println(msg)
	}

	logger("Hello World")
	logger("Hello World")
}
