package main

import "fmt"

func main() {
	/*
		During a typical panic, normal execution stops, all deferred function calls in that goroutine are
		executed, and the program crashes with a log message. This log message includes the panic
		value, which is usually an error message of some sort, and, for each goroutine, a stack trace
		showing the stack of function calls that were active at the time of the panic.

		- Built-in panic function can be used to panic when impossible situation happens
		- When a panic occurs, all deferred functions are run in reverse order, starting with those of the
		topmost function on the stack and proceeding up to main
	*/

	/*
		# Recover
		Giving up is usually the right response to a panic, but not always. It might be possible to
		recover in some way, or at least clean up the mess before quitting . For example, a web server
		that encounters an unexpected problem could close the connection rather than leave the client
		hanging , and during development, it might report the error to the client too.

		If the built-in recover function is called within a deferred function and the function containing the defer statement is panicking , recover ends the current state of panic and returns the
		panic value. The function that was panicking does not continue where it left off but returns
		normally. If recover is called at any other time, it has no effect and returns nil.
	*/

	panicRecover()

}

func panicRecover() {

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	panic("Hii")
}
