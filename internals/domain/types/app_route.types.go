package types

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AppRoute struct {
	Root        *chi.Mux
	Middlewares []func(http.Handler) http.Handler
}
