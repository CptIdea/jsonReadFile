package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJsonFile(path string, out interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file error: %w", err)
	}

	err = json.Unmarshal(bytes, out)
	if err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	return nil
}
