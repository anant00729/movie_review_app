package routes

import (
	"database/sql"

	_c "github.com/anant00729/movie_review_app/controllers"
	"github.com/gorilla/mux"
)

// GetUserRoutes as
func GetUserRoutes(db *sql.DB, r *mux.Router) {
	_c.Connect(db)
	r.HandleFunc("/auth/getAllUsers", _c.GetAllUsers).Methods("GET")
	r.HandleFunc("/movies/getAllMovies", _c.GetMovieData).Methods("GET")
}
