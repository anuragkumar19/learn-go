package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
}

func fetch(url string, wg *sync.WaitGroup) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	res, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "FETCH: %v", err)
	}

	body, err := io.ReadAll(res.Body)

	fmt.Println(string(body))

	wg.Done()
}
