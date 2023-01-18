package handlers

import (
	"encoding/json"
	"net/http"
	g "packages-api/internal"
)

func HandleJSONData(w http.ResponseWriter, r *http.Request) {
	fileName := "json/packages/contrib/amd64/amd64.json"
	jsonData, err := g.GetJSONData(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.MarshalIndent(jsonData, "", "\t")
	if err != nil {
		http.Error(w, "Error marshalling json: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
