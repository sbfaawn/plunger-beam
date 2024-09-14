package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Driver      string `yaml:"driver"`
		Address     string `yaml:"address"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Port        string `yaml:"port"`
		Database    string `yaml:"database"`
		IsPopulated bool   `yaml:"isPopulated"`
		IsMigrate   bool   `yaml:"isMigrate"`
	} `yaml:"database"`
	Cache struct {
		Drive    string `yaml:"drive"`
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		DbNum    int    `yaml:"dbNum"`
		Password string `yaml:"password"`
	} `yaml:"cache"`
}

var cfg *Config

func Read() {
	env := flag.String("env", "development", "Environment (development|production)")
	flag.Parse()

	configPath := filepath.Join("configs", "config."+*env+".yaml")
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

}

func GetConfig() *Config {
	return cfg
}
