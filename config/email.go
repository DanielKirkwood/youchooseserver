package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Email struct {
	SMPT     string
	Port     int
	Username string
	Password string
}

func EmailClient() Email {
	var email Email
	envconfig.MustProcess("EMAIL", &email)

	return email
}
