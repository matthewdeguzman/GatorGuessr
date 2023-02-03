package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var users []User

func initializeRouter() {
	// initialize router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{username}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{username}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{username}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {

	// mock data - @TODO: Implement real database with azure
	users = append(users, User{Username: "madmatt10125", Password: "good_password"})
	users = append(users, User{Username: "thereal_throatgoat", Password: "great_password :o"})

	initializeRouter()
}
