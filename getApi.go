package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
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

type Location struct {
	Id           int
	Locations    []string
	ConcertDates string
}

type Date struct {
	Id    int
	Dates []string
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

//? this function reads the location API from a local file and prints the relevant information
func handleLocation(w http.ResponseWriter, r *http.Request) {
	locations, err := getLocation("./data/locations.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//* Parse the HTML template
	tmpl, err := template.ParseFiles("./view/html/location.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//* Executes the template by passing the location data
	err = tmpl.Execute(w, locations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//* Get the location from the API
func getLocation(filePath string) ([]Location, error) {
	//* Opening the local JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier : %w", err)
	}
	defer file.Close()

	//* Reading the contents of the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier : %w", err)
	}

	//* Decoding JSON content in a Location slice
	var locations []Location
	err = json.Unmarshal(content, &locations)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du décodage du JSON : %w", err)
	}
	return locations, nil
}

//? this function reads the date API from a local file and prints the relevant information
func handleDate(w http.ResponseWriter, r *http.Request) {
	dates, err := getDates("./data/dates.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//* Parse the HTML template
	tmpl, err := template.ParseFiles("./view/html/date.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//* Executes the template by passing the date data
	err = tmpl.Execute(w, dates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//* Get the date from the API
func getDates(filePath string) ([]Date, error) {
	//* Opening the local JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier : %w", err)
	}
	defer file.Close()

	//* Reading the contents of the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier : %w", err)
	}

	//* Decoding JSON content in a date slice
	var dates []Date
	err = json.Unmarshal(content, &dates)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du décodage du JSON : %w", err)
	}
	return dates, nil
}
