package config

import (
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var configurations *Config

type DbConfig struct {
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
	DbSSLMode  string
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWTSecret   string
	DbConfig    *DbConfig
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

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB_HOST environment variable is not set")
		os.Exit(1)
	}

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		fmt.Println("DB_PORT environment variable is not set")
		os.Exit(1)
	}

	dbPort, err := strconv.ParseInt(dbPortStr, 10, 32)
	if err != nil {
		fmt.Println("DB_PORT environment variable is not set or invalid", err)
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB_USER environment variable is not set")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB_PASSWORD environment variable is not set")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB_NAME environment variable is not set")
		os.Exit(1)
	}

	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		fmt.Println("DB_SSL_MODE environment variable is not set")
		os.Exit(1)
	}

	dbConfig:= &DbConfig{
			DbHost:     dbHost,
			DbPort:     int(dbPort),
			DbUser:     dbUser,
			DbPassword: dbPassword,
			DbName:     dbName,
			DbSSLMode:  dbSSLMode,
		}

	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JWTSecret:   jwtSecret,
		DbConfig:    dbConfig,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}