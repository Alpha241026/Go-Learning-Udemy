package main

import (
	"embed"
	"fmt"
	"log"
)

var name = "Joseph"

// special comment below - embeds the entire "public" directory into the executable
//
//go:embed public
var public embed.FS

func main() {

	//read a file from the embedded filesystem
	data, err := public.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	fmt.Println(name)

}
