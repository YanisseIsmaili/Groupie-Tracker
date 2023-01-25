package main

import (
	"fmt"
	"net/http"
	"html/template"
	"path"
	"time"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./view/html")))

	fmt.Println("[INFO] - Starting the server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("[ERROR] - Server could not start properly.\n ", err)
	}
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := "index.html"
	if req.URL.Path == "/" {
		name = "index.html"
	} else {
		name = path.Base(req.URL.Path)
	}

	data := struct {
		Time time.Time
	}{
		Time: time.Now(),
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

