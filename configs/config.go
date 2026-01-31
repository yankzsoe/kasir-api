package configs

import (
	"os"
)

// Config holds all application configuration
type Config struct {
	Database Database
	Server   Server
	API      API
}

// Database configuration
type Database struct {
	Driver string
	Path   string
}

// Server configuration
type Server struct {
	Port string
	Mode string
}

// API configuration
type API struct {
	Version     string
	Title       string
	Description string
	Host        string
	BasePath    string
	Schemes     string
}

// LoadConfig loads all configurations from environment variables
func LoadConfig() *Config {
	return &Config{
		Database: Database{
			Driver: getEnv("DB_DRIVER", "sqlite"),
			Path:   getEnv("DB_PATH", "kasir.db"),
		},
		Server: Server{
			Port: getEnv("PORT", "8080"),
			Mode: getEnv("GIN_MODE", "release"),
		},
		API: API{
			Version:     getEnv("API_VERSION", "v1"),
			Title:       getEnv("API_TITLE", "Kasir API"),
			Description: getEnv("API_DESCRIPTION", "A complete Point of Sale (POS) API"),
			Host:        getEnv("SWAGGER_HOST", "localhost:8080"),
			BasePath:    getEnv("SWAGGER_BASEPATH", "/api/v1"),
			Schemes:     getEnv("SWAGGER_SCHEMES", "http"),
		},
	}
}

// getEnv retrieves environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
