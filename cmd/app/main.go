package main

import (
	"os"
	"smux/config"
	"smux/di/wire"
	"smux/internals/application/server"
	"smux/internals/common/constants/routerconst"
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

	controllerRegistry, err := wire.InitializeControllerRegistry()
	if err != nil {
		exitCode = 1
		return
	}

	// init server
	srv, serverError := wire.InitServer(controllerRegistry, configVars, routerconst.V1)
	if serverError != nil {
		exitCode = 1
		return
	}

	//start the server
	srvErr := server.StartServer(srv)
	if srvErr != nil {
		exitCode = 1
		return
	}
}
