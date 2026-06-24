package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(a, b int) (int, error) { //returning both integer result and error check
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) { //another example of mutli value return function
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]
	return
}

func main() {
	value, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	firstName, lastName := splitName("Arman Rizvi")
	fmt.Println(firstName, lastName)
}
