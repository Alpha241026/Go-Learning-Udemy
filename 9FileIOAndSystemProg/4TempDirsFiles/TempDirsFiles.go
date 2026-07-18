package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tempFile, err := os.CreateTemp("", "logs.txt") //create a temporary file in the system's temp directory
	if err != nil {
		log.Fatal(err)
	}
	defer func() { //clean up the temporary file when the program exits
		fmt.Println("Removing tempFile", tempFile.Name())
		_ = os.Remove(tempFile.Name())
	}()

	_, err = tempFile.Write([]byte("Hello World\n")) //write sample content into the temporary file
	if err != nil {
		log.Fatal(err)
		tempFile.Close()
		return
	}

	tempDir, err := os.MkdirTemp("", "my_app_logs") //create a temporary directory for short-lived data
	if err != nil {
		log.Fatal(err)
	}

	defer func() { //remove the temporary directory before exiting.
		fmt.Println("Removing tempDir", tempDir)
		_ = os.Remove(tempDir)
	}()
}
