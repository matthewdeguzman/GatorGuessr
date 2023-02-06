package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// retrieves the users from the database and encodes it
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// retrieves the params and finds the first user with the id
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
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
	w.Header().Set("Content-Type", "application/json")
	// finds the user with the id and reads the request and updates the user
	params := mux.Vars(r)
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Model(params["id"]).Update(user)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// retrieves user from the database and deletes it
	params := mux.Vars(r)
	var user User
	db.Delete(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
