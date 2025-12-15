package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	Version  string
	HttpPort string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}
	configuration = &Config{
		Version:  version,
		HttpPort: port,
	}
}

func GetConfig() Config {
	if configuration == nil {
		loadConfig()
	}
	return *configuration
}
