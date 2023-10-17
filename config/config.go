package config

import (
	"log"
	"smux/internals/common/constants"
	"sync"

	godotenv "github.com/joho/godotenv"
)

type EnvConfig struct {
	Port string
}

var once sync.Once

func LoadConfig() (*EnvConfig, error) {
	var env *EnvConfig
	once.Do(func() {
		godotenv.Load(".env")
		envMap, err := godotenv.Read()
		if err != nil {
			log.Fatal(1)
		}
		env = &EnvConfig{
			Port: envMap[constants.Port],
		}
	})
	return env, nil
}
