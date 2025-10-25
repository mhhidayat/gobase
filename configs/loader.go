package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return &Config{
		Server: &Server{
			Port: os.Getenv("SERVER_PORT"),
			Host: os.Getenv("SERVER_HOST"),
		},
		Database: &Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
		},
	}
}
