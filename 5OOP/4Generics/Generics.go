package main

import "fmt"

type Number interface { //interface implemeting type constraint by union of selected types only
	int | float32 | float64 | string
}

func Sum[T Number](numbers ...T) T { //implementing generics in variadic function Sum to allow flexibility of types
	var total T
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	v := Sum("Jane", "Mark") //here the Sum function will just concatenate the strings since there are no numbers to add
	fmt.Printf("%T\n", v)
}
