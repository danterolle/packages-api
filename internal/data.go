package internal

import (
	"encoding/json"
	"os"
	c "packages-api/cache"
	p "packages-api/models"
)

func GetJSONData(fileName string) (p.PackageSet, error) {
	var jsonData p.PackageSet

	// Check if the data is already in cache
	if data, ok := c.Get(fileName); ok {
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
	c.Set(fileName, jsonData)

	return jsonData, nil
}
