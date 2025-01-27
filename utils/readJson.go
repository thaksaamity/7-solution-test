package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadJSON(path string) ([][]int, error) {
	// Read the JSON file
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Create a variable to store the data as [][]int
	var data [][]int

	// Unmarshal the JSON data into the variable
	if err := json.Unmarshal(fileData, &data); err != nil {
		log.Printf("Error unmarshalling JSON data: %v", err)
		return nil, err
	}

	return data, nil
}
