package routes

import (
	"backend/handlers"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func LinkRoutes(r *mux.Router) {
	postRepository := repositories.RepositoryLink(mysql.DB)
	h := handlers.HandlerLink(postRepository)

	r.HandleFunc("/", h.CreateLink).Methods("POST")
	r.HandleFunc("/", h.GetLongURL).Methods("GET")
	r.HandleFunc("/{unique_id}", h.GetLink).Methods("GET")
}
