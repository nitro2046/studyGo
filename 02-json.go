package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"kara-,omitempty"`

	Actors []string
}


func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"A", "B"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"C"}},
		{Title: "Bullitt", Year:  1968, Color: true, Actors: []string{"D", "E"}},
	}

	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{
		Title string
		Year int `json:"released"`
	}
	err = json.Unmarshal(data, &titles)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Println(titles)
}
