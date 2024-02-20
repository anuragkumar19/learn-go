package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	res, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("FETCH: %v", err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		ch <- fmt.Sprintf("FETCH READ: %v", err)
	}

	ch <- string(body)
}
