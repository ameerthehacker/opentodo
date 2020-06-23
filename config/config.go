package config

import (
	"log"
	"os"
)

type Config struct {
	Name         string
	Required     bool
	DefaultValue string
}

// func GetConfig can return default value, return value from environment, panic if it was required and no value could be found
func GetConfig(config Config) string {
	envValue := os.Getenv(config.Name)

	if config.Required && len(envValue) == 0 {
		log.Panicf("Config %s is required", config.Name)
	} else if len(envValue) == 0 {
		return config.DefaultValue
	}

	return envValue
}
