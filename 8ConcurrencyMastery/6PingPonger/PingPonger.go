package main

import (
	"context"
	"fmt"
	"time"
)

// continuously sends "ping" messages until the context is cancelled
func ping(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("ping: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

// continuously sends "pong" messages until the context is cancelled
func pong(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("pong: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	//root context with cancellation support
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//shared channel receiving messages from both workers
	pingerCh := make(chan string)
	//notification channel used to tell main that processing finished
	done := make(chan struct{})

	//start concurrent workers
	go ping(ctx, pingerCh)
	go pong(ctx, pingerCh)

	go func() {
		//stop after five seconds
		timeout := time.After(5 * time.Second)
		for {
			//wait until timeout or an incoming message
			select {
			case <-timeout:
				fmt.Println("operation completed")
				close(pingerCh)
				//notify main that work is complete
				done <- struct{}{}
				return

			//receive and print messages
			case msg := <-pingerCh:
				fmt.Println(msg)
			}
		}
	}()

	<-done
	fmt.Println("done")
}
