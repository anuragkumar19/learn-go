package main

import "fmt"

// Functions are reusable piece of code which are generally called more than once from different parts of code
// It mainly consist of three parts -
// 1. params list 2. result list 3. body
// params and results can be named or unnamed
// name of parameter and result don't effect its type
// NOTE: No concept of default parameter

func add(x int, y int) int { return x + y }
func add2(x, y int) int    { return x + y }
func zero(int, int) int    { return 0 }

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", add2)
	fmt.Printf("%T\n", zero)

	// Multiple return
	fmt.Println(someFunc())

	x, err := doSomething(4)

	if err != nil {
		//...
	}

	fmt.Println(x)

	// We can pass multiple return as args to a func with many args
	add(someFunc())
	// similar to
	x, y := someFunc()
	add(x, y)

	// Function values
	// Functions are like other value they have type, can be assigned to other variable, can be passed of returned from a function
	// A function value may be called like any other function, Zero value nil, Calling nil func will panic
	// Not comparable but can be compared with nil
	f := add
	f(2, 3)
}

// Multiple return
func doSomething(x int) (int, error) {
	if x < 0 {
		return 0, fmt.Errorf("Cannot be -ve")
	}

	return x, nil
}

// Multiple returns can be named for documentation purpose but its not required
// If return value are name go initialize the variable to its zero value in same lexical block as params
// Now we can assign to these variable & use "naked return" which basically return to value of the these variables
// It is not advised to use naked return as they don't make code easier to read
func someFunc() (x, y int) {
	// do stuff
	x = 2
	y = 4
	return // This is naked return
}
