package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config holds config for the whole thing.
// It also has some context info for now. We may move it to "Context" type
type Config struct {
	// Slack stuff
	Slacks map[string]*ConfigSlack `yaml:"slack"`

	// Telegram stuff
	Telegrams map[string]*ConfigTelegram `yaml:"telegram"`

	// Mailgun stuff
	Mailguns map[string]*ConfigMailgun `yaml:"mailgun"`

	// Twilio stuff
	Twilios map[string]*ConfigTwilio `yaml:"twilio"`
}

// ConfigSlack is a config block for Slack
type ConfigSlack struct {
	Key string `yaml:"key"`
}

// ConfigTelegram is a config block for Telegram
type ConfigTelegram struct {
	Token  string `yaml:"token"`
	ChatID string `yaml:"chat_id"`
}

// ConfigMailgun is a config block for Mailgun
type ConfigMailgun struct {
	Domain string `yaml:"domain"`
	Key    string `yaml:"key"`
}

// ConfigTwilio is the config for Twilio
type ConfigTwilio struct {
	Sid   string `yaml:"sid"`
	Token string `yaml:"token"`
	From  string `yaml:"from"`
	To    string `yaml:"to"`
}

// NewConfig is a constructor and make a new Config object
// The profile name comes from a command line
func NewConfig(fn string, profileName string) *Config {
	var cfg Config

	rawData, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(rawData), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("read config => %v", cfg)
	log.Printf("read config => %v", cfg.Slacks)

	return &cfg
}
