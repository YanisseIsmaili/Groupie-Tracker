package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/artist.html", handleArtist)
	http.HandleFunc("/location.html", handleLocation)
	fmt.Println("[INFO] server starting at port 8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
