package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var configuration Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}
	httpPort := os.Getenv("SERVER_PORT")
	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}
	config := Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
	}
	configuration = config
}

func GetConfig() Config {
	if configuration.Version == "" {
		LoadConfig()
	}
	return configuration
}
