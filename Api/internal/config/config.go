package config

import "os"

type ConfigDatabase struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func NewConfigDatabase() *ConfigDatabase {
	// Pega o .env
	return &ConfigDatabase{
		Host:     os.Getenv("DATABASE_HOST"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
		Name:     os.Getenv("DATABASE_NAME"),
		Port:     os.Getenv("DATABASE_PORT"),
	}

}
