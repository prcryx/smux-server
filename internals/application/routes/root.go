package routes

import (
	"smux/internals/application/app"
	"smux/internals/application/routes/middlewares"
	"smux/internals/domain/types"

	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(app *app.App, server *types.Server) {
	root := chi.NewRouter()
	root.Use(middlewares.Cors())
	MountAll(root, app.Version, app.ControllerRegistry)
	server.Router = root
}
