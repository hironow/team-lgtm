package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Port               int    `envconfig:"HTTP_PORT" default:"8080"`
	UseDatastore       bool   `envconfig:"USE_DATASTORE" default:"false"`
	DatastoreProjectID string `envconfig:"DATASTORE_PROJECT_ID" default:""`
}

func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, err
	}

	log.Printf("%+v", env)

	return &env, nil
}
