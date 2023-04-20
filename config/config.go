package config

import (
	"log"

	"github.com/joho/godotenv"
)

// A Config is a config object for entire application.
// Can contain multiple configs within for different purposes.
// For example, a Api config can be extracted into it's own struct.'
type Config struct {
	Api      Api
	Database Database
	Email    Email
}

// New returns a new Config struct
func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	return &Config{
		Api:      API(),
		Database: DataStore(),
		Email:    EmailClient(),
	}
}
