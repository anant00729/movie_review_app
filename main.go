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
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/catch_up")
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
