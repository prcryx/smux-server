package types

import (
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Port   int
	Router *chi.Mux
}
