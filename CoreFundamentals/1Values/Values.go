//some basic types of values in Go

package main

import "fmt"

func main() {
	fmt.Printf("Hello" + "World\n") //string (concatenation of two string literals here)

	fmt.Println(1 + 1) //integer (simple arithmetic addition here)

	fmt.Println(3.14) //float

	fmt.Println(true, false) //boolean

	fmt.Printf("%+v\n", []int{1, 2, 3}) //slice (extandable/dynamic array...containing only integers here)

	var t any // The 'any' type: an alias for interface{}, can hold values of any type
	// By default, an uninitialized variable of type 'any' is nil

	fmt.Printf("%+v\n", t)
}
