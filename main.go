package main

import (
	"log"
	"net/http"
	h "packages-api/handlers"
)

// curl --request GET 'http://localhost:8080/packages?branch=main&arch=amd64&package=nginx'
func main() {
	http.HandleFunc("/packages", h.GetPackage)
	http.ListenAndServe(":8080", nil)
	log.Println("Endpoint /packages set, check http://localhost:8080")
	log.Println("Try curl --request GET 'http://localhost:8080/packages?branch=main&arch=amd64&package=nginx'")
}
