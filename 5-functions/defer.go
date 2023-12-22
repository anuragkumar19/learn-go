package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex
var m = make(map[string]int)

func main() {
	/*
		Syntactically, a defer statement is an ordinary function or method call prefixed by the
		keyword defer. The function and argument expressions are evaluated when the statement is
		executed, but the actual call is deferred until the function that contains the defer statement
		has finished, whether normally, by executing a return statement or falling off the end, or
		abnormally, by panicking . Any number of calls may be deferred; they are executed in the
		reverse of the order in which they were deferred.
	*/

	resp, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	f, err := os.Open("")
	if err != nil {
		//...
	}
	defer f.Close()

}

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

// Defer func to print value of double
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

// Modify the result
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
