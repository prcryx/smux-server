package main

import (
	"log"
	"os"
	"smux/config"
)

func main() {
	var exitCode int = 0
	defer func() {
		if exitCode != 0 {

			os.Exit(exitCode)
		}
	}()
	configVars, configError := config.LoadConfig()
	if configError != nil {
		exitCode = 1
		return
	}

	log.Println(configVars)
}
