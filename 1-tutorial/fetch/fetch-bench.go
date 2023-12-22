package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("total: %d ms\n", time.Since(start).Milliseconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

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

	ch <- string(fmt.Sprintf("%s: %d ms :%s", url, time.Since(start).Milliseconds(), body))
}
