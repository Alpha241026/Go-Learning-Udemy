package main

import "fmt"

func main() {
	messages := make(chan string, 3) //buffered channel capable of holding 3 messages

	fmt.Println("Sending messages to buffered channel")

	//these sends dont block because buffer still has space
	messages <- "first message"
	messages <- "second message"
	messages <- "third message"

	//receive messages in FIFO order
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
