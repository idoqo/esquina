package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

func store(router chi.Router) {
	router.Get("/", getAllStores)
	router.Post("/", createNewStore)
	router.Route("/{id}", func (router chi.Router) {
		router.Delete("/", deleteStore)
	})
}

func getAllStores(writer http.ResponseWriter, request *http.Request) {}

func createNewStore(writer http.ResponseWriter, request *http.Request) {}

func deleteStore(writer http.ResponseWriter, request *http.Request) {}
