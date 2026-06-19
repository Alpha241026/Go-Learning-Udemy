package main

import "fmt"

func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	*val = *val * 10 //dereferencing
	fmt.Printf("modifyPointer: %+v\n", &val)
}

func main() {

	age := 10
	agePtr := &age
	fmt.Printf("agePtr: %d\n", agePtr)
	fmt.Printf("Age addr: %d\n\n", &age)

	fmt.Println(age)
	modifyPointer(&age)
	fmt.Println(age) //value changed after calling from function
}
