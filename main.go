package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./view/html")))

	http.HandleFunc("/",IndexHandler)

	fmt.Println("[INFO] - Starting the server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("[ERROR] - Server could not start properly.\n ", err)
	}
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {

}
