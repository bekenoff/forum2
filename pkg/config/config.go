package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Host string `json:"Host"`
	Port string `json:"Port"`
	DB   DB
}

type DB struct {
	Driver   string `json:"driver"`
	Dsn      string `json:"dsn"`
	Database string `json:"database"`
}

func NewConfig() (*Config, error) {
	configFile := "/Users/bekenov/Desktop/forum-main/pkg/config/config.json"
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("ERROR: Read file in config encountered problem: %v", err)
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Printf("ERROR: Unmarshalling in config encountered problem: %v", err)
		return nil, err
	}
	return &config, nil
}
