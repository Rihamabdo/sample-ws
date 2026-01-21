package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser    string
	DBPass    string
	DBHost    string
	DBPort    string
	DBName    string
	JWTSecret string
}

func Load() (Config, error) {
	_ = godotenv.Load()
	c := Config{
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if c.DBUser == "" || c.DBHost == "" || c.DBPort == "" || c.DBName == "" || c.JWTSecret == "" {
		return Config{}, fmt.Errorf("missing env vars (DB_* / JWT_SECRET)")
	}
	return c, nil
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName,
	)
}
