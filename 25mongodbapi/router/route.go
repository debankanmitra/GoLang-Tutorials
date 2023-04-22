package router

import (
	"mdb/controller"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()

	// routing
	router.HandleFunc("/api/movies", controller.GetMymovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkasWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.Deletemovie).Methods("DELETE")

	return router
}
