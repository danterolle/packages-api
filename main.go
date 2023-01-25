package main

import (
	"net/http"
	h "packages-api/handlers"
)

// Example: http://localhost:8080/packages?branch=main&arch=amd64 ~16 mb
// curl --request GET 'http://localhost:8080/packages?branch=main&arch=amd64&package=nginx'
func main() {
	http.HandleFunc("/packages", h.HandleSinglePackage)
	http.ListenAndServe(":8080", nil)
}
