package config

import (
	"os"
	"strconv"
)

func LoadFromEnv() (Config, error) {
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return Config{}, err
	}

	c := Config{
		Port: port,
	}

	return c, nil
}

type Config struct {
	Port int
}
