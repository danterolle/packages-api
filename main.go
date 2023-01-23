package main

import (
	"net/http"
	h "packages-api/handlers"
)

// Example: http://localhost:8080/packages?branch=main&arch=amd64 ~16 mb
func main() {
	http.HandleFunc("/packages", h.HandleJSONData)
	http.ListenAndServe(":8080", nil)
}
