package models

import "database/sql"

// MovieBase (Model)
type MovieBase struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Movies  []Movie `json:"movies"`
}

// Movie (Model)
type Movie struct {
	ID            int    `json:"id"`
	OriginalTitle string `json:"originaltitle"`
	PosterPath    string `json:"posterpath"`
	BackdropPath  string `json:"backdroppath"`
}

// GetAllMovies methods
func GetAllMovies(db *sql.DB) MovieBase {
	var _b MovieBase
	results, err := db.Query("SELECT id , original_title , poster_path, backdrop_path FROM movies;")
	if err != nil {
		_b = MovieBase{Status: false, Message: "Error in Quering the data"}
		return _b
	}

	var movies []Movie

	for results.Next() {
		var m Movie
		error := results.Scan(&m.ID, &m.OriginalTitle, &m.PosterPath, &m.BackdropPath)
		if error != nil {
			_b = MovieBase{Status: false, Message: "Error in Scaning the data"}
			return _b
		}
		movies = append(movies, m)
	}
	_b = MovieBase{Status: false, Message: "", Movies: movies}
	return _b
}
