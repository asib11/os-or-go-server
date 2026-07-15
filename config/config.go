package config

import (
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWTSecret   string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION environment variable is not set")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME environment variable is not set")
		os.Exit(1)
	}

	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		fmt.Println("HTTP_PORT environment variable is not set")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPortStr, 10, 32)
	if err != nil {
		fmt.Println("HTTP_PORT environment variable is not set or invalid", err)
		os.Exit(1)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT_SECRET environment variable is not set")
		os.Exit(1)
	}

	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JWTSecret:   jwtSecret,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}