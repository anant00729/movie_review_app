package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/anant00729/movie_review_app/routes"
	_ "github.com/go-sql-driver/mysql"
)

// ConnectToDB Methods
func ConnectToDB() *sql.DB {
	db, err := sql.Open("mysql", "b6c386390f3b05:2fc2cd6e@tcp(us-cdbr-iron-east-03.cleardb.net:3306)/heroku_a29f673bda288da")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Conneted to DB")
	}
	//defer db.Close()
	return db
}

func init() {

}

func main() {
	db := ConnectToDB()
	r := routes.InitRouter(db)
	log.Fatal(http.ListenAndServe(":5344", r))
}
