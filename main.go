package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("[INFO] - Starting the server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("[ERROR] - Server could not start properly.\n ", err)
	}
}
