// pkg/config/config.go

package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig represents the configuration for the database.
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// LoadDBConfig loads database configuration from .env file.
func LoadDBConfig() (*DBConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %s", err)
	}

	return &DBConfig{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Database: os.Getenv("DATABASE"),
	}, nil
}
