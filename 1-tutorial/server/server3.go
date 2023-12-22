package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s, %s, %s, %s\n", r.Host, r.Method, r.Proto, r.URL.Path)
		fmt.Printf("RemoteAddr: %s\n", r.RemoteAddr)

		fmt.Println("Headers:")
		for k, v := range r.Header {
			fmt.Printf("%s:%s\n", k, v)
		}

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		fmt.Println("Form:")
		for k, v := range r.Form {
			fmt.Printf("%s:%s\n", k, v)
		}
	})

	log.Fatal(http.ListenAndServe(":3454", nil))
}
