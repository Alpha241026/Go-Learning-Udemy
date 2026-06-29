package main

import (
	"fmt"
	"time"
)

type Employee struct { //creating new variable Employee of type struct
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

// value receiver
func (e Employee) FullName() string { //adding method to struct, mimicking OOP
	return e.FirstName + " " + e.LastName
}

func (e *Employee) Deactivate() { //adding method to struct, mimicking OOP...also using pointer to change actual data instead of copy
	e.IsActive = false
}

func NewEmployee(id int, firstName, lastName, position string, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

func main() {
	fmt.Println()
}
