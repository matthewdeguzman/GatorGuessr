package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	ID        uint `gorm:"primarykey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func getPassword() string {
	// Get password from credentials.txt
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(ex)))), "credentials.txt")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	password := ""
	for scanner.Scan() {
		password += scanner.Text()
	}
	password = strings.TrimSpace(password)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return password
}

func initializeMigration() {

	const DB_USERNAME = "cen3031"
	const DB_NAME = "user_database"
	const DB_HOST = "cen3031-server.mysql.database.azure.com"
	const DB_PORT = "3306"
	var password = getPassword()
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
	r.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		switch r.Method {
		case "GET":
			getUsers(w, r)
		case "POST":
			createUser(w, r)
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	r.HandleFunc("/api/users/{username}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		switch r.Method {
		case "GET":
			getUser(w, r)
		case "PUT":
			updateUser(w, r)
		case "DELETE":
			deleteUser(w, r)
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
