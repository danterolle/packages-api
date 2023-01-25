package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"packages-api/internal"
	"packages-api/models"
	"packages-api/utils"
	"path/filepath"
)

func GetPackage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	branch := vars.Get("branch")
	arch := vars.Get("arch")
	packageName := vars.Get("package")
	if branch == "" || arch == "" || packageName == "" {
		http.Error(w, "branch, architecture and package name are required", http.StatusBadRequest)
		return
	}

	branch = utils.SanitizeInput(branch)
	arch = utils.SanitizeInput(arch)
	packageName = utils.SanitizeInput(packageName)

	if !utils.CheckAllowlist(branch, models.Branch) {
		log.Println("Invalid branch")
		return
	}
	if !utils.CheckAllowlist(arch, models.Arch) {
		log.Println("Invalid architecture")
		return
	}

	fileName := filepath.Join("json", "packages", branch, arch, arch+".json")
	jsonData, err := internal.GetJSONData(fileName)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var packageData interface{}
	for _, p := range jsonData.Packages {
		if p.Name == packageName {
			packageData = p
			break
		}
	}

	jsonBytes := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(jsonBytes)
	jsonEncoder.SetIndent("", "  ")

	err = jsonEncoder.Encode(packageData)
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes.Bytes())
}
