package internal

import (
	config "server/internal"
)

type AWSConfig struct {
	AWS struct {
		AWSAccessKeyID     string `json:"aws_access_key_id"`
		AWSSecretAccessKey string `json:"aws_secret_access_key"`
		AWSRegion          string `json:"aws_region"`
	} `json:"AWS"`
}


func ReadAWSConfig(filename string) (*AWSConfig, error) {
	var awsConfig AWSConfig
	err := config.ReadConfig(filename, &awsConfig)
	if err != nil {
		return nil, err
	}
	return &awsConfig, nil
}