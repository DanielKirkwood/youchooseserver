package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// A Api gets and stores the API env variables
// required for the rest API. Standard defaults
// provided if env variables not set.
type Api struct {
	Name              string        `default:"youchoose_api"`
	Host              string        `default:"0.0.0.0"`
	Port              string        `default:"3080"`
	ReadHeaderTimeout time.Duration `split_words:"true" default:"60s"`
	GracefulTimeout   time.Duration `split_words:"true" default:"8s"`

	RequestLog bool `split_words:"true" default:"false"`
}

func API() Api {
	var api Api
	envconfig.MustProcess("API", &api)

	return api
}
