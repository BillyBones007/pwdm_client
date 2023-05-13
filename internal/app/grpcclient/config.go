package grpcclient

import (
	"encoding/json"
	"os"
)

const configFile string = "config.json"

// Config - configuration the client.
type Config struct {
	ServerAddr string `json:"server_address"`
}

// setFileConfig - set configuration from config file.
func setFileConfig() (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(configFile)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		return config, err
	}

	return config, nil
}
