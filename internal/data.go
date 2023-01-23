package internal

import (
	"encoding/json"
	"os"
	p "packages-api/models"
)

func GetJSONData(fileName string) (p.PackageSet, error) {
	var jsonData p.PackageSet
	var err error
	var file *os.File

	ch := make(chan error, 2)
	defer close(ch)

	go func() {
		file, err = os.Open(fileName)
		ch <- err
	}()

	go func() {
		defer file.Close()
		err = json.NewDecoder(file).Decode(&jsonData)
		ch <- err
	}()

	for i := 0; i < 2; i++ {
		if err := <-ch; err != nil {
			return p.PackageSet{}, err
		}
	}

	return jsonData, nil
}
