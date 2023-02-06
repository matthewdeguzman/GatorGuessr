package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func initializeMigration() {

	// Build connection string
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	} else {
		fmt.Println("Server Connected Successfully")
	}

	// migrates the server if necessary
	db.AutoMigrate(&User{})

}

func initializeRouter() {
	// initialize router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{username}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{username}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{username}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {

	initializeMigration()
	initializeRouter()
}
