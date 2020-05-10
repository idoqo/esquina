package handler

import (
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"net/http"
)

// NewHandler creates a new API request handler based on github.com/go-chi/chi
func NewHandler() http.Handler{
	router := chi.NewRouter()
	router.Route("/api", func (router chi.Router) {

	})
	return cors.AllowAll().Handler(router)
}
