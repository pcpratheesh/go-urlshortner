package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/pcpratheesh/go-urlshortner/models"
)

func LoadConfigurations(tag string) (models.EnvVariables, error) {
	//load the configurations
	var env models.EnvVariables
	err := envconfig.Process(tag, &env)

	return env, err
}
