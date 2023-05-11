package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./view/html")))
	http.HandleFunc("/artist.html", handleArtist)
	http.HandleFunc("/location.html", handleLocation)
	http.HandleFunc("/date.html", handleDate)
	fmt.Println("[INFO] server starting at port 8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
