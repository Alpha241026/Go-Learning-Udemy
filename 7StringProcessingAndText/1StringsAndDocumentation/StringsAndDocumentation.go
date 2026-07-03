package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abc"
	s2 := strings.Clone(s1) //clone creates a copy of the string

	fmt.Println(s1, s2)

	b := strings.Builder{} //efficiently constructs strings without repeated allocations
	b.WriteString("Here is an example")

	fmt.Println(b.String())

	//converts string to lowercase / uppercase
	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s1))

	s3 := "   test sssssss     "
	fmt.Println("s3", len(s3))
	s3 = strings.TrimSpace(s3) //removes leading and trailing whitespace
	fmt.Println("s3 after trim", len(s3))

	fmt.Println(strings.HasSuffix("test@gmail.com", "gmail.com"))     //checks whether a string ends with the given suffix
	fmt.Println(strings.HasPrefix("test@gmail.com", "test"))          //checks whether a string starts with the given prefix
	fmt.Println(strings.Replace("test@gmail.com", "test", "john", 1)) //replaces occurrences of one substring with another

	parts := strings.Split("test@gmail.com", "@") //splits a string using a separator
	username, domain := parts[0], parts[1]
	fmt.Println(username, domain)

	parts = strings.Fields("jane example") //splits a string on whitespace
	username, domain = parts[0], parts[1]
	fmt.Println(username, domain)

	fmt.Println(strings.Join(parts, ",")) //joins multiple strings using a separator.
}
