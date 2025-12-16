package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	if err := write(c); err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err = os.WriteFile(configPath, jsonData, 0o660); err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := homeDir + "/" + configFileName

	return configDir, nil
}
