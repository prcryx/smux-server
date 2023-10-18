package config

import (
	"log"
	"smux/internals/common/constants"
	"smux/internals/common/utils"
	"sync"

	godotenv "github.com/joho/godotenv"
)

type ConfigVars struct {
	Port int
}

var once sync.Once

func LoadConfig() (*ConfigVars, error) {
	var env *ConfigVars
	once.Do(func() {
		godotenv.Load(".env")
		envMap, err := godotenv.Read()
		if err != nil {
			log.Fatal(1)
		}
		env = &ConfigVars{
			Port: utils.StringToInt(envMap[constants.Port]),
		}
	})
	return env, nil
}
