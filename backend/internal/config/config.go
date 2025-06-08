package config

import (
    "os"
    "strconv"
)

type Config struct {
    Port            string
    DatabaseURL     string
    RedisURL        string
    JWTSecret       string
    GoogleClientID  string
    GoogleClientSecret string
    Environment     string
}

func Load() *Config {
    return &Config{
        Port:            getEnv("PORT", "8080"),
        DatabaseURL:     getEnv("DATABASE_URL", ""),
        RedisURL:        getEnv("REDIS_URL", "redis://localhost:6379"),
        JWTSecret:       getEnv("JWT_SECRET", "your-secret-key"),
        GoogleClientID:  getEnv("GOOGLE_CLIENT_ID", ""),
        GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
        Environment:     getEnv("ENVIRONMENT", "development"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}