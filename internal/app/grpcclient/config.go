package grpcclient

import (
	"encoding/json"
	"os"

	"github.com/BillyBones007/pwdm_client/internal/storage"
)

const configFile string = "config.json"

// Config - configuration the client.
type Config struct {
	ServerAddr string `json:"server_address"`
	Storage    *storage.Storage
}

// setConfig - set configuration from config file.
func setConfig() (*Config, error) {
	config := &Config{}
	config.Storage = storage.NewStorage()
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
