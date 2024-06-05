package internal

import (
	"encoding/json"
	"os"
)

func ReadConfig(filename string, config interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return err
	}

	return nil
}