// Package config XiXiOrchard/config/config.go
package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	APIKey     string `yaml:"api_key"`
	DBHost     string `yaml:"db_host"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
	LogLevel   string `yaml:"log_level"`
	MarketAPI  string `yaml:"market_api"`
}

var Cfg Config

// LoadConfig reads the config.yaml file and unmarshals it into the Config struct.
func LoadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&Cfg)
	if err != nil {
		return err
	}

	log.Printf("Configuration loaded successfully from %s", configPath)
	return nil
}
