package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	DatabaseURL string
	RedisURL    string
	ImageDir    string
	BaseURL     string
	Version     string
}

func Load() *Config {
	return &Config{
		Port:        getEnvInt("PORT", 8080),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://martyria:martyria@localhost:5432/martyria?sslmode=disable"),
		RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
		ImageDir:    getEnv("IMAGE_DIR", "./data/images"),
		BaseURL:     getEnv("BASE_URL", "http://localhost:8080"),
		Version:     getEnv("VERSION", "0.1.0"),
	}
}

func (c *Config) DBConnString() string {
	return c.DatabaseURL
}

func (c *Config) Addr() string {
	return fmt.Sprintf(":%d", c.Port)
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return fallback
}
