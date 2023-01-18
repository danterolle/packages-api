package internal

import (
	"encoding/json"
	"os"
	p "packages-api/models"
)

func GetJSONData(fileName string) (p.PackageSet, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return p.PackageSet{}, err
	}
	defer file.Close()

	var jsonData p.PackageSet
	if err := json.NewDecoder(file).Decode(&jsonData); err != nil {
		return p.PackageSet{}, err
	}

	return jsonData, nil
}
