package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with a capacity of 2
	ch := make(chan int, 2)

	// Start a goroutine to send numbers
	go func() {
		ch <- 1   // Send 1 to the channel
		ch <- 2   // Send 2 to the channel
		close(ch) // Close the channel
	}()

	// Receive numbers from the channel
	for num := range ch {
		fmt.Println("Received:", num)
	}
}
