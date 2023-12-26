package config

import (
	"1c/model"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// loadConfig загружает конфигурацию из файла YAML.
func LoadConfig(filename string) (*model.Config, error) {
	const op = "1c.Config.LoadConfig"
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	var config model.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &config, nil
}

// SaveConfig сохраняет конфигурацию в YAML.
func SaveConfig(filename string, config model.Config) error {
	const op = "1c.Config.SaveConfig"

	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}
