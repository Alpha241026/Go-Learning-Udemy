package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct { //field names used when encoding to JSON or XML
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone"`
	IsActive bool   `json:"active" xml:"active"`
}

func main() {
	jane := user{
		Name:     "Jane",
		Age:      20,
		IsActive: true,
		Phone:    "123-456-7890",
	}

	byteSlice, err := json.Marshal(jane) //convert the Go struct into JSON
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteSlice)) //convert the JSON bytes into a printable string
}
