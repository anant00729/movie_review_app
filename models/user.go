package models

import (
	"database/sql"
)

// Base MODEL
type Base struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Users   []User `json:"users"`
}

// User MODEL
type User struct {
	//Status   bool   `json:"status"`
	//Message  string `json:"message"`
	Username   string `json:"username"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	ProfileImg string `json:"profileImg"`
}

// GetAllUsers methods
func GetAllUsers(db *sql.DB) Base {
	var _b Base
	results, err := db.Query("SELECT username, email, mobile, profile_img FROM users;")
	if err != nil {
		_b = Base{Status: false, Message: "Error in Quering the data"}
		return _b
	}

	var users []User

	for results.Next() {
		var m User
		error := results.Scan(&m.Username, &m.Email, &m.Phone, &m.ProfileImg)
		if error != nil {
			_b = Base{Status: false, Message: "Error in Scaning the data"}
			return _b
		}
		users = append(users, m)
	}
	_b = Base{Status: false, Message: "", Users: users}
	return _b
}
