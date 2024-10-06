package main

import (
	"fmt"
	"talentforce/calculator"
	// Adjust the import path as needed
)

func main() {
	a := 10.0
	b := 5.0

	fmt.Println("Addition:", calculator.Add(a, b))
	fmt.Println("Subtraction:", calculator.Subtract(a, b))
	fmt.Println("Multiplication:", calculator.Multiply(a, b))
	fmt.Println("Division:", calculator.Divide(a, b))
	fmt.Println("Division by zero:", calculator.Divide(a, 0))
}
