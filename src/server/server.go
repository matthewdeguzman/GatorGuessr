package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/matthewdeguzman/GatorGuessr/src/server/credentials"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const DB_USERNAME = "cen3031"
const DB_NAME = "user_database"
const DB_HOST = "cen3031-server.mysql.database.azure.com"
const DB_PORT = "3306"
const DSN = DB_USERNAME + ":" + credentials.DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

type User struct {
	ID        uint `gorm:"primarykey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
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
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
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
