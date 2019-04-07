package controllers

import (
	"encoding/json"
	"net/http"

	_m "github.com/anant00729/movie_review_app/models"
)

// var db *sql.DB

// // Connect get connection Object
// func Connect(_db *sql.DB) {
// 	db = _db
// }

// GetAllUsers as
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var _b _m.Base
	_b = _m.GetAllUsers(db)
	json.NewEncoder(w).Encode(_b)
}
