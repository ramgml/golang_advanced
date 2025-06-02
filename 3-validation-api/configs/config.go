package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email string
	Password string
	Address string
	Urn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		Email: os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Address: os.Getenv("ADDRESS"),
		Urn: "email.json",
	}
}
