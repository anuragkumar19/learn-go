package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)

	log.Fatal(http.ListenAndServe(":3454", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "COUNT: %d\n", count)
	mu.Unlock()
}
