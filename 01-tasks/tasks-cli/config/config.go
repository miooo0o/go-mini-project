package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	FilePath string  `json:"file_path"`
	Alarms   []Alarm `json:"alarms"`
}

type Alarm struct {
	Time    string `json:"time"`
	Message string `json:"message"`
}

func LoadConfig(configPath string) (Config, error) {
	var cfg Config

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return cfg, fmt.Errorf("config file not found: %s", configPath)
	} else if err != nil {
		return cfg, fmt.Errorf("failed to access config file: %w", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("ailed to parse config file (%s): %w", configPath, err)
	}
	return cfg, nil
}
