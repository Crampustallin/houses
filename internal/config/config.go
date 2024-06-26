package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func (c *Config) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func NewConfig() *Config {
	return &Config{
		DBHost: getenv("POSTGRES_HOST", "localhost"), 
		DBPort: getenv("POSTGRES_PORT", "5432"),
		DBUser: getenv("POSTGRES_USER", "postgres"),
		DBPassword: getenv("POSTGRES_PASSWORD", "postgres"),
		DBName: getenv("POSTGRES_DB", "rams_db"),
	}
}

func getenv(key, defaultV string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultV
	}
	return value
}
