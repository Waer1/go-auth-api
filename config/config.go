package config

import (
	"log"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
)

type EnvFile struct {
	AppPort     string        `valid:"required,numeric"`
	DbUser      string        `valid:"required"`
	DbPassword  string        `valid:"required"`
	DbHost      string        `valid:"required,host"`
	DbPort      string        `valid:"required,numeric"`
	DbName      string        `valid:"required"`
	JwtSecret   string        `valid:"required,length(32|64)"`
	JwtExpireIn time.Duration `valid:"required"`
}

var Config EnvFile

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Map environment variables to the Config struct
	Config = EnvFile{
		AppPort:     getEnv("APP_PORT"),
		DbUser:      getEnv("DB_USER"),
		DbPassword:  getEnv("DB_PASSWORD"),
		DbHost:      getEnv("DB_HOST"),
		DbPort:      getEnv("DB_PORT"),
		DbName:      getEnv("DB_NAME"),
		JwtSecret:   getEnv("JWT_SECRET"),
		JwtExpireIn: parseDuration("JwtExpireIn"), // Parse the expiration time as a time.Duration
	}

	// Validate the configuration
	if _, err := govalidator.ValidateStruct(Config); err != nil {
		log.Fatalf("Configuration validation error: %v", err)
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func parseDuration(key string) time.Duration {
	valueStr := getEnv(key)
	value, err := time.ParseDuration(valueStr)
	if err != nil {
		log.Fatalf("Error parsing %s as duration: %v", key, err)
	}
	return value
}
