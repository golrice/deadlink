package utils

import (
	"encoding/json"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func LoadJsonConfig(filePath string, config interface{}) error {
	fileData, err := ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, config)
}

func SaveJsonData(filePath string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return WriteFile(filePath, jsonData)
}
