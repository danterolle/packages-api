package internal

import (
	"encoding/json"
	"fmt"
	"os"
	p "packages-api/models"
	"sync"
)

func GetJSONData(fileName string) (p.PackageSet, error) {
	var jsonData p.PackageSet
	var err error
	var file *os.File

	ch := make(chan error, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return p.PackageSet{}, fmt.Errorf("file %s does not exist", fileName)
	}

	go func() {
		defer wg.Done()
		file, err = os.Open(fileName)
		ch <- err
	}()

	go func() {
		defer wg.Done()
		defer file.Close()
		err = json.NewDecoder(file).Decode(&jsonData)
		ch <- err
	}()

	wg.Wait()
	for i := 0; i < 2; i++ {
		if err := <-ch; err != nil {
			return p.PackageSet{}, err
		}
	}

	return jsonData, nil
}

/*
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
*/
