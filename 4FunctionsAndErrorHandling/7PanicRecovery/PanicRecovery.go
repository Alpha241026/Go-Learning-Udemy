package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong in mightPanic") //crashes the program...not to be used as an exception handling mechanism
	}
	fmt.Println("this function executed without a panic")
}

func recoverable() { //function called when a panic is about to occur
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()
	mightPanic(true)
}

func main() {
	mightPanic(false)
	mightPanic(true)
}
