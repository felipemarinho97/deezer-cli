package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultFormat string `json:"default_format"`
	DefaultLimit  int    `json:"default_limit"`
	CacheEnabled  bool   `json:"cache_enabled"`
	CacheTTL      int    `json:"cache_ttl_seconds"`
}

var defaultConfig = Config{
	DefaultFormat: "table",
	DefaultLimit:  25,
	CacheEnabled:  true,
	CacheTTL:      300,
}

func Load() (*Config, error) {
	configPath := getConfigPath()
	
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &defaultConfig, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return &defaultConfig, fmt.Errorf("failed to read config: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return &defaultConfig, fmt.Errorf("failed to parse config: %w", err)
	}

	if config.DefaultFormat == "" {
		config.DefaultFormat = defaultConfig.DefaultFormat
	}
	if config.DefaultLimit == 0 {
		config.DefaultLimit = defaultConfig.DefaultLimit
	}

	return &config, nil
}

func Save(config *Config) error {
	configPath := getConfigPath()
	configDir := filepath.Dir(configPath)

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

func getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ".deezer-cli-config.json"
	}

	return filepath.Join(homeDir, ".config", "deezer-cli", "config.json")
}