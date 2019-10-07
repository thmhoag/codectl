package template

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func LoadValuesFromFile(pathToFile string) (map[string]interface{}, error) {
	fileBytes, err := getFileBytes(pathToFile)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}
	// Try and load json
	if err := json.Unmarshal(fileBytes, &result); err != nil {
		return nil, err
	}

	containsItems := len(result) > 0
	if containsItems {
		// we're done, send it back
		return result, nil
	}

	// Try and load yaml
	if err := yaml.Unmarshal(fileBytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func getFileBytes(pathToManifest string) ([]byte, error) {
	file, err := os.Open(pathToManifest)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ioutil.ReadAll(file)
}
