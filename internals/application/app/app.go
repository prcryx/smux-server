package app

import (
	"smux/config"
	"smux/di/container"
	"smux/internals/domain/valobj"
)

type App struct {
	ControllerRegistry *container.ControllerRegistry
	ConfigVars         *config.ConfigVars
	Version            valobj.AppVersion
}

func NewApp(controllerRegistry *container.ControllerRegistry, configVars *config.ConfigVars, version string) *App {
	return &App{
		ControllerRegistry: controllerRegistry,
		ConfigVars:         configVars,
		Version:            valobj.GetAppVersion(version),
	}
}
