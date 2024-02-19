package main

import "fmt"

func main() {
	// Variadic function can be called with varying number of arguments, eg- fmt.Printf
	// To declare a variadic function, type of final parameter is preceded by an ellipsis("...")
	// This mean this func can be called with any number of args for this type
	fmt.Println(sum(1, 2, 3, 4, 5))

	// to invoke a variadic function when the arguments are already in a slice: place an
	// ellipsis after the final argument.
	ints := []int{1, 2, 3, 4}
	sum(ints...)

	fmt.Printf("%T\n", sum)
}

func sum(val ...int) int {
	// type of val :- []int
	total := 0
	for _, x := range val {
		total += x
	}
	return total
}
