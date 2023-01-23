package handlers

import (
	"encoding/json"
	"net/http"
	g "packages-api/internal"
)

func HandleJSONData(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	branch := vars.Get("branch")
	arch := vars.Get("arch")
	fileName := "json/packages/" + branch + "/" + arch + "/" + arch + ".json"

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
