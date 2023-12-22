package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

func main() {
	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	// Go struct to json
	b, err := json.Marshal(movies)

	if err != nil {
		fmt.Println(err)
	}

	var movies2 []Movie

	// JSON to go struct
	json.Unmarshal(b, &movies2)

	fmt.Printf("%#v\n", movies2)

	// Prettify JSON
	b2, _ := json.MarshalIndent(movies, "", "    ") // ignoring err
	fmt.Println(string(b2))

	var movie Movie

	err = json.Unmarshal([]byte(`{}`), &movie)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(movie)
}
