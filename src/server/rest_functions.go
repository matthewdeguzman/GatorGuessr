package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

// getUsers returns all the users from the database
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// getUser returns a specified user from the database
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var user User
	db.First(&user, "Username = ?", params["username"])
	json.NewEncoder(w).Encode(user)
}

// createUser creates a new user and inserts into the database
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// updateUser updates a user with the sent information
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user, "Username = ?", params["username"])
	if user.Username != "" {
		json.NewDecoder(r.Body).Decode(&user)
		db.Save(&user)
	}
	json.NewEncoder(w).Encode(user)
}

// deleteUser deletes a user from the database
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.Delete(&user, "Username = ?", params["username"])
	json.NewEncoder(w).Encode(user)
}
