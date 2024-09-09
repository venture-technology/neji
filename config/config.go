package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name     string  `yaml:"name"`
	Server   Server  `yaml:"server"`
	Discord  Discord `yaml:"discord"`
	Database Database
}

type Server struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Secret string `yaml:"string"`
}

type Discord struct {
	Token string `yaml:"token"`
}

type Database struct {
	User     string `yaml:"dbuser"`
	Port     string `yaml:"dbport"`
	Host     string `yaml:"dbhost"`
	Password string `yaml:"dbpassword"`
	Name     string `yaml:"dbname"`
	Schema   string `yaml:"schema"`
}

var config *Config

func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	config = &conf
	return config, nil
}

func Get() *Config {

	// if was created to run tests
	if config == nil {
		config, err := Load("../../../config/config.yaml")
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
		return config
	}

	return config
}
