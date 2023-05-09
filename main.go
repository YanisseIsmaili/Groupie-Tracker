package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/artists", handleArtist)
	http.HandleFunc("/locations", handleLocation)
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
