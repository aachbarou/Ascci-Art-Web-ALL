package main

import (
	"fmt"
	"net/http"

	Handler "groupie-tracker/Handlers"
)

func main() {
	// Process the artists data (e.g., print to console, save to a file, etc.)

	fmt.Println("server  start at  Port 8080 : http://localhost:8080")
	// http.HandleFunc("/", Handler.Handle_Home)
	http.HandleFunc("/", Handler.HAndleHOme)
	// http.HandleFunc("/download", Handler.DownloadHandler)
	http.ListenAndServe(":8080", nil)
}
