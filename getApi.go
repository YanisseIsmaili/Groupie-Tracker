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

type Location struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Data struct {
	Artists   []Artist
	Locations []Location
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	artists, err := getArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Artists = artists
	locations, err := getLocations("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Locations = locations
	tmpl, err := template.ParseFiles("./view/html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleArtist(w http.ResponseWriter, r *http.Request) {
	artists, err := getArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("./view/html/artist.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getArtists(filePath string) ([]Artist, error) {
	resp, err := http.Get(filePath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier : %w", err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier : %w", err)
	}
	var artists []Artist
	json.Unmarshal(content, &artists)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du décodage du JSON : %w", err)
	}
	return artists, nil
}

func handleLocation(w http.ResponseWriter, r *http.Request) {
    locations, err := getLocations("https://groupietrackers.herokuapp.com/api/locations")
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

    //* Executes the template by passing the locations' data
    tmpl.Execute(w, locations)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func getLocations(filePath string) ([]Location, error) {
    resp, err := http.Get(filePath)
    if err != nil {
        return nil, fmt.Errorf("erreur lors de l'ouverture du fichier : %w", err)
    }
    defer resp.Body.Close()

    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("erreur lors de la lecture du fichier : %w", err)
    }

    var locResp LocationResponse
    err = json.Unmarshal(content, &locResp)
    if err != nil {
        return nil, fmt.Errorf("erreur lors du décodage du JSON : %w", err)
    }

    var locations []Location
    for _, l := range locResp.Index {
        loc := Location{
            ID:        l.ID,
            Locations: l.Locations,
            Dates:     l.Dates,
        }
        locations = append(locations, loc)
    }

    return locations, nil
}
