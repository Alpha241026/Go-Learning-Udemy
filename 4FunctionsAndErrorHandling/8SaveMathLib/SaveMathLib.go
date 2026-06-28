package main

import (
	"fmt"
	"strings"
)

type MathError struct { //custom error type
	Operation string
	InputA    int
	InputB    int
	Message   string
}

const divErrMsg = "division by zero is not allowed"

func (e MathError) Error() string { //error interface
	var inputs []string
	if e.Operation == "Division" {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}
	return fmt.Sprintf("Math error in %s (%s): %s",
		e.Operation,
		strings.Join(inputs, ","),
		e.Message)
}

func Sum(numbers ...int) int { //for caluclating adding numbers
	defer fmt.Println("Sum finished")

	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func SafeDivision(a, b int) (int, error) { //for dividing numbers
	if b == 0 {
		return 0, &MathError{
			Operation: "Division",
			InputA:    a,
			InputB:    b,
			Message:   divErrMsg,
		}
	}
	return a / b, nil
}

func main() {
	fmt.Println(Sum(1, 2, 3))

	value, err := SafeDivision(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SafeDivision", value)
}
