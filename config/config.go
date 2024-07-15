package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var GlobalConfig *Config

type Config struct {
	BotToken      string `yaml:"bot_token"`
	NewsAPIToken  string `yaml:"newsapi_token"`
	Proxy         struct {
		Enabled bool   `yaml:"enabled"`
		URL     string `yaml:"url"`
	} `yaml:"proxy"`
}

func ReadConfig() {
	yamlFile, err := os.ReadFile("./config/config.yaml")

	if err != nil {
		log.Fatalf("Could not read config file: %v", err)
		return
	}

	if err := yaml.Unmarshal(yamlFile, &GlobalConfig); err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
		return
	}
}
