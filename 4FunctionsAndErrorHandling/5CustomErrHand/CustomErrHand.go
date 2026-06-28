package main

import (
	"errors"
	"fmt"
	"time"
)

// predefined error variables
var ErrDivisionByZero = errors.New("divison by zero")
var ErrNumTooLarge = errors.New("number too large")

type OpError struct { //custom error type
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func (op OpError) Error() string { //implementing error interface for OpError
	return op.Message
}

func NewOpError(op string, code int, message string, t time.Time) *OpError { //returns pointer to new OpError instance
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

func DoSomething() error { //placeholder function to simulate failing operation and return custom OpError
	return NewOpError("doSomething", 100, "do something failed", time.Now())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	if a > 1000 {
		return 0, ErrNumTooLarge
	}
	return a / b, nil
}

func main() {
	value, err := divide(1001, 1)
	if err != nil {
		if errors.Is(err, ErrDivisionByZero) { //check if error is due to division by zero
			fmt.Println("divide by zero")
		} else if errors.Is(err, ErrNumTooLarge) { //check if error is due to exceeding allowed limit
			fmt.Println("number too large")
		}
		return //exit if error ocurred
	}
	fmt.Println(value) //print if no error
}
