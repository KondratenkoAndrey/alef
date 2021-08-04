package config

import "os"

type Config struct {
	Environment string
	DB          *DBConfig
}

type DBConfig struct {
	Host     string
	Username string
	Password string
	DbName   string
}

func GetConfig() *Config {
	return &Config{
		Environment: os.Getenv("APP_ENV"),
		DB: &DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
		},
	}
}
