package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config holds the configuration for all messaging services.
type Config struct {
	Slacks    map[string]*Slack    `yaml:"slack"`
	Telegrams map[string]*Telegram `yaml:"telegram"`
	Mailguns  map[string]*Mailgun  `yaml:"mailgun"`
	Twilios   map[string]*Twilio   `yaml:"twilio"`
}

// Slack is the configuration for a Slack webhook.
type Slack struct {
	Key string `yaml:"key"`
}

// Telegram is the configuration for a Telegram bot.
type Telegram struct {
	Token  string `yaml:"token"`
	ChatID string `yaml:"chat_id"`
}

// Mailgun is the configuration for Mailgun email service.
type Mailgun struct {
	Domain string `yaml:"domain"`
	Key    string `yaml:"key"`
}

// Twilio is the configuration for Twilio SMS service.
type Twilio struct {
	Sid   string `yaml:"sid"`
	Token string `yaml:"token"`
	From  string `yaml:"from"`
	To    string `yaml:"to"`
}

// Load reads and parses the configuration file.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	return &cfg, nil
}

// FindConfigPath locates the configuration file in standard locations.
func FindConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting home directory: %w", err)
	}

	paths := []string{
		filepath.Join(homeDir, ".config", ".msgr.conf"),
		filepath.Join(homeDir, ".msgr.conf"),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	return "", fmt.Errorf("config file not found in: %v", paths)
}
