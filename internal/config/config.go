package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration loaded from environment.
type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	// AutoMigrate runs GORM AutoMigrate on startup. Enable in dev for fast
	// iteration; disable in production and rely on SQL migrations instead.
	AutoMigrate bool
	// JWTSecret signs auth tokens. JWTExpiryHours is the token lifetime.
	JWTSecret      string
	JWTExpiryHours int
}

// Load reads configuration from a .env file (if present) and the environment.
func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "postgres"),
		DBName:      getEnv("DB_NAME", "pets"),
		DBSSLMode:   getEnv("DB_SSLMODE", "disable"),
		AutoMigrate:    getEnv("AUTO_MIGRATE", "true") == "true",
		JWTSecret:      getEnv("JWT_SECRET", "dev-secret-change-me"),
		JWTExpiryHours: getEnvInt("JWT_EXPIRY_HOURS", 24),
	}
}

// DSN builds the PostgreSQL connection string.
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}
