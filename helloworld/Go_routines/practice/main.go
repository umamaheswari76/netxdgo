package main

import (
	"fmt"
)

func sendData(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i // Send data into the channel
	}
	close(ch) // Close the channel to signal that no more data will be sent
}

func receiveData(ch <-chan int, done chan<- bool) {
	for {
		value, ok := <-ch // Receive data from the channel
		if !ok {
			// Channel is closed, so we're done
			break
		}
		fmt.Printf("Received: %d\n", value)
	}
	done <- true // Signal that we're done
}

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	// Create a done channel for synchronization
	done := make(chan bool)

	// Start a goroutine to send data into the channel
	go sendData(ch)

	// Start a goroutine to receive data from the channel
	go receiveData(ch, done)

	// Use a select statement to wait for both goroutines to finish
	select {
	case <-done:
		fmt.Println("Data sending and receiving completed.")
	}

	// Buffered Channel Example
	bufferedCh := make(chan int, 2) // Create a buffered channel with capacity 2

	// Send data into the buffered channel (won't block until capacity is exceeded)
	bufferedCh <- 1
	bufferedCh <- 2

	// Receiving data from the buffered channel
	fmt.Printf("Buffered Channel: %d %d\n", <-bufferedCh, <-bufferedCh)

	// Close the buffered channel
	close(bufferedCh)
}