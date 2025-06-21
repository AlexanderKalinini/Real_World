package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const MigrationsPath = "./internal/database/migrations"

func init() {
	var err = godotenv.Load("../.env")
	if err != nil {
		fmt.Println("could not load .env:", err)
	}
}

type Config struct {
	App   AppConfig
	MySql MySqlConfig
}

type AppConfig struct {
	MigrationsPath string
}

type MySqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func LoadConfig() *Config {
	return &Config{
		App: AppConfig{
			MigrationsPath: "./internal/database/migrations",
		},
		MySql: MySqlConfig{
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
		},
	}
}
