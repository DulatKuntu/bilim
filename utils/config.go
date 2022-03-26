package utils

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	Port       string `yaml:"port"`
	StaticPort string `yaml:"static_port"`
	Dburi      string `yaml:"dburi"`
	Dbname     string `yaml:"dbname"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	dburi, exists := os.LookupEnv("MONGO_HOST")

	if !exists {
		dburi = "mongodb://localhost:27017"
	}
	log.Print(dburi)
	StaticPortHTTP, exists := os.LookupEnv("StaticPortHTTP")

	if !exists {
		return nil, errors.New(".env variables not set")
	}
	APIPortHTTP, exists := os.LookupEnv("APIPortHTTP")

	if !exists {
		return nil, errors.New(".env variables not set")
	}

	config.Dburi = dburi
	config.StaticPort = StaticPortHTTP
	config.Port = APIPortHTTP
	config.Dbname = "evision"

	return config, nil
}
