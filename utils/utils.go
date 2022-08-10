package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJsonFile(path string, out interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open file error: %w", err)
	}

	err = json.NewDecoder(file).Decode(out)
	if err != nil {
		return fmt.Errorf("decode json error: %w", err)
	}

	return nil
}
