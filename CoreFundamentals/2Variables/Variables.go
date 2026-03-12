//basics of variables in go

package main

import "fmt"

func main() {
	var greeting string //zero value for string is an empty ""
	greeting = "Salutations!"

	fmt.Println(greeting)

	var count int //zero value for int is 0
	count = 10
	fmt.Println(count)

	var isRunning bool //zero value for bool is false
	isRunning = true

	fmt.Println(isRunning)

	var firstName, lastName string //multiple variables on same line
	firstName = "John"
	lastName = "Doe"
	fmt.Println(firstName, lastName)

	email := "abc@email.com" //type inference
	fmt.Println(email)
}
