package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Version string

type Config struct {
	Port int
	Env  string
}

// Load reads the values of .env file
func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	env := os.Getenv("APP_ENV")
	return Config{
		Port: port,
		Env:  env,
	}
}
