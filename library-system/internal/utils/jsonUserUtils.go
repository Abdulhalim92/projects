package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"projects/internal/model"
)

func ReadJsonFromFileUser(filename string, m *map[int]model.User) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to Open File: %w", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading from file: %w", err)
	}
	if len(data) == 0 {
		return nil
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		return fmt.Errorf("error unmarshaling file: %w", err)
	}
	return nil
}
func WriteJsonToUserFile(filename string, users *map[int]model.User) error {
	data, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to decode data: %w", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error to write json to the user file: %w", err)
	}
	return nil
}
