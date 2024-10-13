package main

import (
	"fmt"
)

func main() {
	// Create an unbuffered channel for the result
	ch := make(chan int)

	// Numbers to add
	num1 := 5
	num2 := 10
	num3 := 15

	// Start a goroutine to perform the addition
	go func() {
		sum := num1 + num2 + num3 // Perform addition
		ch <- sum                 // Send result to channel
	}()
	result := <-ch
	fmt.Println("The sum of ", num1, num2, num3, "is :", result)
}
