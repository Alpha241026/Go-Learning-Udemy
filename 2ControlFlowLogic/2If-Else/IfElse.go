package main

import "fmt"

func main() {
	tmp := 25

	//if else
	if tmp > 30 {
		fmt.Println("Greater than 30")
	} else {
		fmt.Println("Less than 30")
	}

	score := 85

	//if else ladder
	if score >= 90 {
		fmt.Println("Grade : A")
	} else if score >= 80 {
		fmt.Println("Grade : B")
	} else if score >= 70 {
		fmt.Println("Grade : C")
	} else {
		fmt.Println("Failed")
	}

	//string to boolean map
	userAccess := map[string]bool{
		"Jane": true,
		"John": false,
	}

	if hasAccess, ok := userAccess["Jane"]; ok && hasAccess { //checking is value and key exist; access not granted if tested for John instead
		fmt.Println("Jane can access the system")
	} else {
		fmt.Println("Access not granted")
	}
}
