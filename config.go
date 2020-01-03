package main

import "io/ioutil"
import "gopkg.in/yaml.v2"
import "log"

// Config holds
type Config struct {
	// Stuff from a config file

	// Slack stuff
	Slacks		map[string]*ConfigSlack		`yaml:"slack"`

	// Telegram stuff
	Telegrams	map[string]*ConfigTelegram	`yaml:"telegram"`

	// Runtime settings. Not to complicate things too much
	ProfileName		string
}

type ConfigSlack struct {
	Key	string	`yaml:"key"`
}

type ConfigTelegram struct {
	Token	string	`yaml:"token"`
	ChatID	string	`yaml:"chat_id"`
}

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

	cfg.ProfileName = profileName

	return &cfg
}
