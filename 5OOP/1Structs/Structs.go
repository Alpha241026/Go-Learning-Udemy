package main

import (
	"fmt"
	"time"
)

type Employee struct { //custom data type Employee with struct
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func main() {
	Jane := Employee{ //instance of Employee where we assign values to each field
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    1000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	fmt.Println(Jane)
}
