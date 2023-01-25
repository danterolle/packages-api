package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"packages-api/internal"
	"packages-api/models"
	"packages-api/utils"
	"path/filepath"
)

func HandleSinglePackage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	branch := vars.Get("branch")
	arch := vars.Get("arch")
	packageName := vars.Get("package")
	if branch == "" || arch == "" {
		http.Error(w, "branch and arch are required", http.StatusBadRequest)
		return
	}

	branch = utils.SanitizeInput(branch)
	arch = utils.SanitizeInput(arch)
	packageName = utils.SanitizeInput(packageName)

	if !utils.CheckAllowlist(branch, models.Branch) {
		fmt.Println("Invalid branch")
		return
	}
	if !utils.CheckAllowlist(arch, models.Arch) {
		fmt.Println("Invalid architecture")
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

	jsonBytes, err := json.MarshalIndent(packageData, "", "\t")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error marshalling json: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
