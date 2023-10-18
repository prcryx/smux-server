//go:build wireinject
// +build wireinject

package wire

import (
	"smux/config"
	"smux/di/container"
	"smux/internals/application/app"
	"smux/internals/application/server"
	"smux/internals/domain/types"

	"github.com/google/wire"
)

// initialize db
// func InitializeDB() *DB{}

func InitializeControllerRegistry() (*container.ControllerRegistry, error) {
	wire.Build(
		container.NewControllerRegistry,
		DataSourceSet,
		RepositorySet,
		UseCaseSet,
		ControllerSet,
	)

	return nil, nil
}

//Intialize Server

func InitServer(controllerRegistry *container.ControllerRegistry, configVars *config.ConfigVars, version string) (*types.Server, error) {
	wire.Build(
		server.NewServer,
		app.NewApp,
	)

	return nil, nil
}
