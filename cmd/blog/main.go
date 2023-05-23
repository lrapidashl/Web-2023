package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

const (
	port         = "localhost:3000"
	dbDriverName = "mysql"
)

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	dbx := sqlx.NewDb(db, dbDriverName)

	mux := mux.NewRouter()
	mux.HandleFunc("/home", index(dbx))
	mux.HandleFunc("/post/{postID}", post(dbx))
	mux.HandleFunc("/admin", admin)
	mux.HandleFunc("/api/post", createPost(dbx)).Methods(http.MethodPost)
	mux.HandleFunc("/api/auth", auth(dbx)).Methods(http.MethodPost)
	mux.HandleFunc("/login", login)

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Start server" + port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() (*sql.DB, error) {
	return sql.Open(dbDriverName, "root:Gjhjkmrfrfirf1_@tcp(localhost:3306)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
}
