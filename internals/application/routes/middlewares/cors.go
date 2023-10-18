package middlewares

import (
	"net/http"
	"smux/internals/common/constants"

	"github.com/go-chi/cors"
)

func Cors() func(http.Handler) http.Handler {
	return cors.Handler(
		cors.Options{
			AllowedOrigins:   constants.AllowedOrigins(),
			AllowedHeaders:   constants.AllowedHeaders(),
			AllowedMethods:   constants.AllowedMethods(),
			AllowCredentials: constants.AllowCredentials,
			MaxAge:           constants.MaxAge,
			ExposedHeaders:   constants.ExposedHeaders(),
		},
	)
}
