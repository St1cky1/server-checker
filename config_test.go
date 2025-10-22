package main

import (
	"testing"
	"time"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("config.yaml")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Проверяем значения
	if config.Interval != 10*time.Second {
		t.Errorf("Expected interval 10s, got %v", config.Interval)
	}

	if config.Request_timeout != 2*time.Second {
		t.Errorf("Expected request timeout 2s, got %v", config.Request_timeout)
	}

	if config.Workers_count != 3 {
		t.Errorf("Expected workers count 3, gpt: %v", config.Workers_count)
	}

	if len(config.Urls) == 0 {
		t.Errorf("Expected at least one URL in config")
	}

	// Проверяем первый URL

	firstUrl := config.Urls[0]

	if firstUrl.Name != "Google" {
		t.Errorf("Expected first URL name 'Google, got '%s''", firstUrl.Name)
	}

	t.Log("Config loading successfully")
}
