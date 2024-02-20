package main

import (
	"fmt"
)

func main() {
	// Primary error handling in go - error as values
	// func which may cause error should return a extra result (mostly last one) for errors

	// For simple cases where only one cause of err is possible we return an boolean generally called ok from the func
	v, ok := GetValue("")

	if !ok {
		//...
	}
	fmt.Println(v)

	// For func that with more cause we use built in `error` interface
	// `error` can nil or non-nil

	/*
		# Error handling strategies
		1. Propagate the error, so that caller can handle it
			- For better error handling we may have to add additional info before propagating the error
			- Error message should not be capitalized or new line should be avoided
			- Be deliberate and consistence so error can be handled in same way, we may have to report the parameter passed to the func with error message
		2. Retry with delay, for limited number of times before giving up
		3. If progress is impossible we should print err message and shutdown program gracefully. Ex- DB failed to connect, we cannot continue without db. So we retry few times then exit the program
		4. Print the error and continue
		5. Ignore the error (_ for error return value)
	*/

	// 1.
	// if err != nil {
	// return 0, fmt.Errorf("extra info: %v", err)
	// }

	// 3.
	// fmt.Println(err); os.Exit(1) or
	//  log.Fatal() or
	//  log.Fatalf()

	// Some err are different than others they should also be dealt with differently
	// io.EOF is returned by many `i/o` func when there is nothing more to read. Its not actually a program intervention so it should be checked
}

func GetValue(key string) (value string, ok bool) {
	if key == "" {
		return "", false
	}

	return "Val", true
}
