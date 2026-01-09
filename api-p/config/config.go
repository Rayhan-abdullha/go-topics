package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	Version    string
	HttpPort   string
	SecretJWT  string
	Sha256Algo string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required!")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}
	secretJwt := os.Getenv("SECRET_JWT")
	if secretJwt == "" {
		fmt.Println("Jwt secret is required!")
		os.Exit(1)
	}
	sha256algo := os.Getenv("SHA256")
	if sha256algo == "" {
		fmt.Println("sha256 algo is required!")
		os.Exit(1)
	}
	config := Config{
		Version:    version,
		HttpPort:   httpPort,
		SecretJWT:  secretJwt,
		Sha256Algo: sha256algo,
	}
	configuration = &config
	return configuration
}

func GetConfig() *Config {
	if configuration == nil {
		LoadConfig()
	}
	return configuration
}
