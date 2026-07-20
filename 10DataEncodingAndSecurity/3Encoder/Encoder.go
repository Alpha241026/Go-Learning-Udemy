package main

import (
	"encoding/json"
	"log"
	"os"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	Password string `json:"-" xml:"-"`
	IsActive bool   `json:"active" xml:"active"`
}

func main() {
	u := user{
		Name:  "John Smith",
		Age:   45,
		Phone: "13812345678",
	}

	//create a JSON encoder that writes directly to standard output
	enc := json.NewEncoder(os.Stdout)
	//encode the struct and stream it to the destination
	if err := enc.Encode(&u); err != nil {
		log.Fatal(err)
	}
}
