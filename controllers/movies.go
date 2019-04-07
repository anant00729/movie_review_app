package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_m "github.com/anant00729/movie_review_app/models"
)

var db *sql.DB

// Connect get connection Object
func Connect(_db *sql.DB) {
	db = _db
}

// GetMovieData method
func GetMovieData(w http.ResponseWriter, r *http.Request) {
	var _b _m.MovieBase
	_b = _m.GetAllMovies(db)
	json.NewEncoder(w).Encode(_b)
}
