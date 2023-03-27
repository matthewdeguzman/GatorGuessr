package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	ID           uint `gorm:"primarykey"`
	Username     string
	Password     string
	DailyScore   uint
	WeeklyScore  uint
	MonthlyScore uint
	TotalScore   uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func initializeMigration() {

	const DB_USERNAME = "cen3031"
	const DB_NAME = "user_database"
	const DB_HOST = "cen3031-server.mysql.database.azure.com"
	const DB_PORT = "3306"
	var password = os.Getenv("DB_PASSWORD");
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
	db.AutoMigrate(&User{})
}

func initializeRouter() {
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		EnableCors(w)
		switch r.Method {
		case "GET":
			GetUsers(w, r)
		case "POST":
			CreateUser(w, r)
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	r.HandleFunc("/api/users/{username}/", func(w http.ResponseWriter, r *http.Request) {
		EnableCors(w)
		switch r.Method {
		case "GET":
			GetUser(w, r)
		case "PUT":
			UpdateUser(w, r)
		case "DELETE":
			DeleteUser(w, r)
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	r.HandleFunc("/api/login/", func(w http.ResponseWriter, r *http.Request) {
		EnableCors(w)
		switch r.Method {
		case "POST":
			ValidateUser(w, r)
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
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
