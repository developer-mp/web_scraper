package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AWS struct {
		AWSAccessKeyID     string `json:"aws_access_key_id"`
		AWSSecretAccessKey string `json:"aws_secret_access_key"`
		AWSRegion          string `json:"aws_region"`
	} `json:"AWS"`
}

func ReadAppConfig(filename string) (*AppConfig, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var appConfig  AppConfig
    err = json.Unmarshal(data, &appConfig)
    if err != nil {
        return nil, err
    }

    return &appConfig, nil
}