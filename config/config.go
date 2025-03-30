package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config содержит конфигурацию приложения
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	ServerPort string
	CookieDomain string
	CookieLifetime int
}

// LoadConfig загружает конфигурацию из .env файла или переменных окружения
func LoadConfig() (*Config, error) {
	// Загрузка .env файла, если он существует
	_ = godotenv.Load()

	config := &Config{
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", ""),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		JWTSecret:  getEnv("JWT_SECRET", ""),
		ServerPort: getEnv("SERVER_PORT", ""),
		CookieDomain: getEnv("COOKIE_DOMAIN", ""),
	}

	cookieLifetime, err := strconv.Atoi(getEnv("COOKIE_LIFETIME", "3600"))
	if err != nil {
		return nil, err
	}
	config.CookieLifetime = cookieLifetime

	return config, nil
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
