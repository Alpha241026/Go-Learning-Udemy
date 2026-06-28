package main

import (
	"fmt"
	"os"
)

// demonstrating how defer statements are executed in the end...and among themselves in reverse order (last in, first out)
func lifoSimpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: deferred")
	defer fmt.Println("Function simpleDefer: deferred")
	fmt.Println("Function simpleDefer: Middle")
}

func main() {
	file, err := os.Create("./defer.text") //creating file defer.text
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //closing file when main ends

	//simpleDefer()
	lifoSimpleDefer()

	fmt.Println("Last in main()")
}
