package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
}

func main() {
	messages := make(chan string) //channel for sending string messages

	users := make(chan user) //channel for sending user structs

	//goroutine sending a string through a channel
	go func() {
		fmt.Println("Sending a message to messages channel")
		messages <- "Hello from messages channel"
	}()

	//goroutine sending a user struct through a channel
	go func() {
		fmt.Println("Sending a message to messages channel")
		users <- user{
			name: "Gopher",
		}
	}()

	//temporary delay to let goroutines start
	time.Sleep(1 * time.Second)

	fmt.Println("About to get message from channel")

	//receive one string from the messages channel
	msg := <-messages
	fmt.Println(msg)

	//receive one user from the users channel
	u := <-users
	fmt.Println(u)
}
