package main

import "fmt"

// not performant but ok
func factorial(x int) int {
	if x == 1 {
		return x
	}

	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(1), factorial(2), factorial(3), factorial(4), factorial(5))
}
