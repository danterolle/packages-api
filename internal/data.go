package internal

import (
	"encoding/json"
	"os"
	p "packages-api/models"
	"packages-api/utils/cache"
)

// GetJSONData is designed to retrieve JSON data from a file and cache the data in memory
// so that it can be reused without having to read the file again.
func GetJSONData(fileName string) (p.PackageSet, error) {
	var jsonData p.PackageSet

	// Check if the data is already in cache
	if data, ok := cache.Get(fileName); ok {
		jsonData = data.(p.PackageSet)
		return jsonData, nil
	}

	file, err := os.Open(fileName)
	if err != nil {
		return p.PackageSet{}, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&jsonData); err != nil {
		return p.PackageSet{}, err
	}

	// Add the data to the cache
	cache.Set(fileName, jsonData)

	return jsonData, nil
}
