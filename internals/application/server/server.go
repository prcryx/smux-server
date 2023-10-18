package server

import (
	"fmt"
	"log"
	"net/http"
	"smux/internals/application/app"
	"smux/internals/application/routes"
	"smux/internals/domain/types"
)

func NewServer(app *app.App) (*types.Server, error) {
	server := initServer(app.ConfigVars.Port)
	routes.SetUpRoutes(app, server)
	return server, nil
}

func initServer(port int) *types.Server {
	return &types.Server{
		Port: port,
	}
}

// start a server
func StartServer(server *types.Server) error {
	log.Printf("server is listening on: %v", server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", server.Port), server.Router); err != nil {
		log.Printf("Server failed: %v", err)
		return err
	}
	return nil
}
