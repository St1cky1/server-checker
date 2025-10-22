package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3" // импортировали yaml библиотеку
)

type Config struct { // Структура Config куда распарсиваем yaml
	Interval        time.Duration `yaml:"interval"`
	Request_timeout time.Duration `yaml:"request_timeout"`
	Workers_count   int           `yaml:"workers_count"`
	Urls            []UrlConfig   `yaml:"urls"`
}
type UrlConfig struct {
	Name           string `yaml:"name"`
	Url            string `yaml:"url"`
	ExpectedStatus int    `yaml:"expected_status"`
}

// Функция, которая распарсивает наш yaml в структуры
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config // Перемменная config будет указателем Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("faild to parse config: %w", err)
	}

	return &config, nil
}
