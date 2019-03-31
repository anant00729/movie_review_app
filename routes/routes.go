package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// InitRouter this is  (METHODS)
func InitRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	GetUserRoutes(db, r)
	//r.HandleFunc("/api/getAllMovies", getMovies).Methods("GET")
	return r
}
