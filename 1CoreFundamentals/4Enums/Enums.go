//Enums (enumerations) implemented using const and iota keywords

package main

import "fmt"

const (
	Sunday = iota //initialized as 0 by defualt, next values auto-incremented by default
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
