package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/idoqo/esquina/esquina"
	"github.com/rs/cors"
	"net/http"
)

// NewHandler creates a new API request handler based on github.com/go-chi/chi
func NewHandler() http.Handler{
	router := chi.NewRouter()
	router.NotFound(notFoundHandler)
	router.MethodNotAllowed(methodNotAllowedHandler)

	router.Route("/api", func (router chi.Router) {
		router.Route("/store", store)
	})
	return cors.AllowAll().Handler(router)
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(404)
	render.Render(writer, request, esquina.ErrorNotFound)
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(405)
	render.Render(writer, request, esquina.ErrorMethodNotAllowed)
}
