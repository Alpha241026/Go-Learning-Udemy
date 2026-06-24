package main

import "fmt"

func sum(numbers ...int) int { //variadic function - can accept 0 or more integers
	total := 0
	for _, number := range numbers { //looping over slice
		total += number
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
}
