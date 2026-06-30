package main

import (
	"fmt"
)

type Person interface { //implementing a contract/interface
	GetName() string
}

type Employee struct { //representing an employee by a struct type
	ID   int
	Name string
}

type BusinessPerson struct { //representing a business person by a struct type
	ID   int
	Name string
}

func (e Employee) GetName() string { //allows Employee to satisfy Person interface
	return e.Name
}

func (e BusinessPerson) GetName() string { //allows BusinessPerson to satisfy Person interface
	return e.Name
}

func displayPerson(p Person) { //depending only the Person interface; working with any type that can provide a name
	fmt.Println(p.GetName())
}

func main() {
	joe := Employee{ //creating Employee instance
		ID:   1,
		Name: "Joe",
	}

	jane := BusinessPerson{ //creating BusinessPerson instance
		ID:   1,
		Name: "Joe",
	}

	//same function works for both as they satisfy Person interface
	displayPerson(joe)
	displayPerson(jane)
}
