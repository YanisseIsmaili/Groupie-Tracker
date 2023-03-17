package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

// ? this function reads the artists API from website and prints the relevant information
func handleArtist(w http.ResponseWriter, r *http.Request) {
	artists, err := getArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//* Parse the HTML template
	tmpl, err := template.ParseFiles("./view/html/artist.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//* Executes the template by passing the artists' data
	tmpl.Execute(w, artists)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// * Get the artists from the API
func getArtists(filePath string) ([]Artist, error) {
	//* Opening the JSON page online
	resp, err := http.Get(filePath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier : %w", err)
	}
	defer resp.Body.Close()

	//* Reading the contents of the link
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier : %w", err)
	}

	//* Decoding JSON content in an Artist slice
	var artists []Artist
	json.Unmarshal(content, &artists)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du décodage du JSON : %w", err)
	}
	return artists, nil
}
