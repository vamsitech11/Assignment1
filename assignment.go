package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Prime numbers between 1 and 10:")
	for i := 1; i <= 10; i++ {
		isPrime := true

		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break

			}

		}
		if isPrime {
			sum += i
			fmt.Println(i)

		}

	}
	fmt.Println(sum)

}
