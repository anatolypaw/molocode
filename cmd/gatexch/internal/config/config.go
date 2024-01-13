package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Server структура для конфигурации сервера.
type Server struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

// Config структура для хранения конфигурационных данных из YAML.
type Config struct {
	MainExchangemarks   Server `yaml:"main_exchangemarks"`
	ReserveExhangemarks Server `yaml:"reserve_exchangemarks"`
	Hub                 Server `yaml:"hub"`
	GetPeriod           int    `yaml:"get_period"` // период запроса кодов в секундах
}

// loadConfig загружает конфигурацию из файла YAML.
func LoadConfig(filename string) (*Config, error) {
	const op = "1c.Config.LoadConfig"
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &config, nil
}

// SaveConfig сохраняет конфигурацию в YAML.
func SaveConfig(filename string, config Config) error {
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
