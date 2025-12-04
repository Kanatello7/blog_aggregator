package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_URL   string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	full_path := filepath.Join(home, configFileName)
	return full_path, nil
}

func Read() (Config, error) {
	config_path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(config_path)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func write(cfg Config) error {
	encoded, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	config_path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(config_path, encoded, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.Username = username
	return write(*cfg)
}
