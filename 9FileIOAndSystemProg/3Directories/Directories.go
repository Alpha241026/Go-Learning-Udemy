package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	dir := "Downloads/static/images"
	//create the directory structure if it doesn't already exist
	if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
		log.Fatal(err)
	}

	//remove the directory and everything inside it
	if err := os.RemoveAll("Downloads"); err != nil {
		log.Fatal(err)
	}

}
