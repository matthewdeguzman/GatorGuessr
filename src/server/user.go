package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user, params["username"])
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// reads the request and puts it into user
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	// appends the new user to the database and encodes it into w
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// @TODO - Ethan
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// @TODO - Ethan
}
