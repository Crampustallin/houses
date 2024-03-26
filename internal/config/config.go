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
		DBHost: getenv("DBHost", "hostname"), 
		DBPort: getenv("DBHost", "hostname"),
		DBUser: getenv("DBHost", "hostname"),
		DBPassword: getenv("DBHost", "hostname"),
		DBName: getenv("DBHost", "hostname"),
	}
}

func getenv(key, defaultV string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultV
	}
	return value
}
