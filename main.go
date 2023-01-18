package main

import (
	"net/http"
	h "packages-api/handlers"
)

func main() {
	http.HandleFunc("/json/", h.HandleJSONData)
	http.ListenAndServe(":8000", nil)
}
