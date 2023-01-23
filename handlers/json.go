package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	c "packages-api/cache"
	g "packages-api/internal"
)

func HandleJSONData(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	branch := vars.Get("branch")
	arch := vars.Get("arch")
	if branch == "" || arch == "" {
		http.Error(w, "branch and arch are required", http.StatusBadRequest)
		return
	}

	if val, ok := c.Get(branch + arch); ok {
		jsonBytes, err := json.MarshalIndent(val, "", "\t")
		if err != nil {
			http.Error(w, "Error marshalling json: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
		return
	}

	fileName := "json/packages/" + branch + "/" + arch + "/" + arch + ".json"
	jsonData, err := g.GetJSONData(fileName)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	c.Set(branch+arch, jsonData)
	jsonBytes, err := json.MarshalIndent(jsonData, "", "\t")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error marshalling json: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
