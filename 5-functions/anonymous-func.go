package main

import "fmt"

func main() {
	// Named functions can be declared only at the package level, but we can use a function literal to
	// denote a function value within any expression.
	// Function literal is same as func declaration but without name. It is an expression & its value is called anonymous function
	a := func() {
		//...
	}

	a()

	// Function declared in this way have access to entire lexical env it is within
	// This mean function also have state
	// This technique is called closure
	f := squares()
	fmt.Println(f(), f(), f(), f())

	//	When anonymous function requires recursion, we must first declare a
	// variable, and then assign the anonymous function to that variable. If these two steps been
	// combined in the declaration, the function literal would not be within the scope of the variable, will cause compile error

	// b := func ()  {
	// 	b() // compile err
	// }

	var b func()

	b = func() {
		b() //ok
	}
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
