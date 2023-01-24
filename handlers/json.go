package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	g "packages-api/internal"
	"packages-api/models"
	"packages-api/utils"
	"packages-api/utils/cache"
	"path/filepath"
)

// HandleJSONData This function is a HTTP handler that is designed to handle requests for JSON data.
func HandleJSONData(w http.ResponseWriter, r *http.Request) {

	// Get the "branch" and "arch" variables from the URL query
	vars := r.URL.Query()
	branch := vars.Get("branch")
	arch := vars.Get("arch")
	// Check if "branch" and "arch" are provided, return an error if they are not
	if branch == "" || arch == "" {
		http.Error(w, "branch and arch are required", http.StatusBadRequest)
		return
	}

	// Check if we have a cached version of the JSON data
	if val, ok := cache.Get(branch + arch); ok {
		// If we have a cached version, convert it to JSON and write it to the response
		jsonBytes, err := json.MarshalIndent(val, "", "\t")
		if err != nil {
			http.Error(w, "Error marshalling json: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
		return
	}

	// Sanitize the "branch" and "arch" variables
	branch = utils.SanitizeInput(branch)
	arch = utils.SanitizeInput(arch)

	if !utils.CheckWhitelist(branch, models.Branch) {
		fmt.Println("Invalid branch")
		return
	}
	if !utils.CheckWhitelist(arch, models.Arch) {
		fmt.Println("Invalid architecture")
		return
	}

	// Construct the file name
	fileName := filepath.Join("json", "packages", branch, arch, arch+".json")
	// Read the JSON data from the file
	jsonData, err := g.GetJSONData(fileName)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Cache the JSON data
	cache.Set(branch+arch, jsonData)
	// Convert the JSON data to bytes and write it to the response
	jsonBytes, err := json.MarshalIndent(jsonData, "", "\t")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error marshalling json: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
