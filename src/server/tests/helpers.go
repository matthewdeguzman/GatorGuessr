package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	endpoints "github.com/matthewdeguzman/GatorGuessr/src/server/endpoints"
	u "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func testInitMigration(t *testing.T) (db *gorm.DB) {
	const DB_USERNAME = "cen3031"
	const DB_NAME = "user_database"
	const DB_HOST = "cen3031-server.mysql.database.azure.com"
	const DB_PORT = "3306"

	// Build connection string
	DSN := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	// migrates the server if necessary
	db.AutoMigrate(&u.User{})

	return db
}

func cleanDB(user *u.User, username string, t *testing.T) {
	db := testInitMigration(t)
	db.Delete(user, "Username = ?", user.Username)
}

func mockGetUsers(w http.ResponseWriter, r *http.Request, t *testing.T) {
	db := testInitMigration(t)
	endpoints.GetUsers(w, r, db)
}

func mockGetUser(w http.ResponseWriter, r *http.Request, username string, t *testing.T) {
	db := testInitMigration(t)

	endpoints.SetHeader(w)

	var user u.User
	endpoints.FetchUser(db, &user, username)

	if user.Username == "" {
		endpoints.WriteErr(w, http.StatusNotFound, "")
	}
	endpoints.EncodeUser(user, w)

}

func mockCreateUser(w http.ResponseWriter, r *http.Request, user u.User, db *gorm.DB, t *testing.T) {
	endpoints.SetHeader(w)
	if endpoints.UserExists(db, user.Username) || user.ID != 0 || user.Password == "" {
		endpoints.WriteErr(w, http.StatusBadRequest, "")
		return
	}

	hash, err := endpoints.EncodePassword(user.Password)

	if err != nil {
		endpoints.WriteErr(w, http.StatusInternalServerError, "")
	}
	user.Password = hash

	db.Create(&user)
	endpoints.EncodeUser(user, w)
}

func getUserTest(username string, t *testing.T) (status int) {
	req, err := http.NewRequest("GET", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockGetUser(w, r, username, t)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
}

func createUserTest(user u.User, t *testing.T) (status int) {
	db := testInitMigration(t)
	req, err := http.NewRequest("POST", "/api/users/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockCreateUser(w, r, user, db, t)
	})

	handler.ServeHTTP(rr, req)

	return rr.Code
}
