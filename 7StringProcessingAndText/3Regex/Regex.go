package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text1 := "Hello world! Welcome to Go"

	regGo, err := regexp.Compile(`Go`) // compiles a regex pattern that searches for "Go"
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Text '%s', matches 'Go': %t\n", text1, regGo.MatchString(text1)) //checks whether the text contains the pattern

	text2 := "Product codes: P123, X342, P789"

	//compiles a regex for product codes beginning with 'P' followed by one or more digits
	rProductP := regexp.MustCompile(`P\d+`)

	//finds the first matching product code
	firstProduct := rProductP.FindString(text2)
	fmt.Println(firstProduct)

	//finds every matching product code
	allPProducts := rProductP.FindAllString(text2, -1)
	fmt.Printf("%+v\n", allPProducts)
}
