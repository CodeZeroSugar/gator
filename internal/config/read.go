package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	configDir, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(configDir)
	if err != nil {
		return Config{}, err
	}

	var configData Config
	if err := json.Unmarshal(data, &configData); err != nil {
		return Config{}, err
	}

	return configData, nil
}
