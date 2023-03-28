package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/matthewdeguzman/GatorGuessr/src/server/endpoints"
	db_user "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func initializeMigration() {

	const DB_USERNAME = "cen3031"
	const DB_NAME = "user_database"
	const DB_HOST = "cen3031-server.mysql.database.azure.com"
	const DB_PORT = "3306"
	var password = os.Getenv("DB_PASSWORD")
	// Build connection string
	DSN := DB_USERNAME + ":" + password + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	} else {
		fmt.Println("Server Connected Successfully")
	}

	// migrates the server if necessary
	db.AutoMigrate(&db_user.User{})
}

func initializeRouter() {
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		endpoints.EnableCors(w)
		switch r.Method {
		case "GET":
			endpoints.GetUsers(w, r, db)
		case "POST":
			endpoints.CreateUser(w, r, db)
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	r.HandleFunc("/api/users/{username}/", func(w http.ResponseWriter, r *http.Request) {
		endpoints.EnableCors(w)
		switch r.Method {
		case "GET":
			endpoints.GetUser(w, r, db)
		case "PUT":
			endpoints.UpdateUser(w, r, db)
		case "DELETE":
			endpoints.DeleteUser(w, r, db)
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	r.HandleFunc("/api/login/", func(w http.ResponseWriter, r *http.Request) {
		endpoints.EnableCors(w)
		switch r.Method {
		case "POST":
			endpoints.ValidateUser(w, r, db)
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	// r.HandleFunc("/api/leaderboard/", func(w http.ResponseWriter, r *http.Request) {
	// 	endpoints.EnableCors(w)
	// 	switch r.Method {
	// 		case
	// 	}
	// })
	r.HandleFunc("/api/leaderboard/{limit}/", func(w http.ResponseWriter, r *http.Request) {
		endpoints.EnableCors(w)
		switch r.Method {
		case "GET":
			endpoints.GetTopUsers(w, r, db)
		case "OPTIONS":
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	initializeMigration()
	initializeRouter()
}
