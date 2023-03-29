package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func addUser(user *u.User, t *testing.T) {
	db := testInitMigration(t)
	db.Create(user)
}

/// MOCK FUNCTIONS ///

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

func mockUpdateUser(w http.ResponseWriter, r *http.Request, userMap map[string]string, username string, db *gorm.DB, t *testing.T) {
	endpoints.SetHeader(w)

	var oldUser u.User
	var updatedUser u.User
	endpoints.FetchUser(db, &oldUser, username)
	endpoints.FetchUser(db, &updatedUser, username)

	if oldUser.Username == "" {
		endpoints.UserDNErr(w)
		return
	}

	// decode user
	for key, element := range userMap {
		switch key {
		case "Username":
			updatedUser.Username = element
		case "ID":
			id, _ := strconv.Atoi(element)
			updatedUser.ID = uint(id)
		case "Score":
			score, _ := strconv.Atoi(element)
			updatedUser.Score = uint(score)

		}
	}
	if oldUser.ID != updatedUser.ID {
		endpoints.WriteErr(w, http.StatusMethodNotAllowed, "405 - Cannot change ID")
		return
	}

	hash, err := endpoints.EncodePassword(updatedUser.Password)
	if err != nil {
		endpoints.HashErr(w)
		return
	}
	updatedUser.Password = hash
	updatedUser.CreatedAt = oldUser.CreatedAt

	db.Save(&updatedUser)
	endpoints.EncodeUser(updatedUser, w)
}

func mockDeleteUser(w http.ResponseWriter, r *http.Request, username string, db *gorm.DB, t *testing.T) {
	endpoints.SetHeader(w)
	var user u.User

	endpoints.FetchUser(db, &user, username)
	if user.Username == "" {
		endpoints.UserDNErr(w)
		return
	}
	db.Delete(&user, "Username = ?", username)
	endpoints.EncodeUser(user, w)
}

func mockGetTopUsers(w http.ResponseWriter, r *http.Request, limit string, db *gorm.DB, t *testing.T) {
	var users []u.User

	lim, err := strconv.Atoi(limit)
	if err != nil {
		endpoints.WriteErr(w, http.StatusBadRequest, "400 - Could not process limit")
		return
	}
	if lim <= 0 {
		endpoints.WriteErr(w, http.StatusBadRequest, "400 - Limit must be a positive integer")
		return
	}

	db.Limit(lim).Order("score desc").Find(&users)
	endpoints.EncodeUsers(users, w)
}

// TESTING FUNCTIONS //

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

func updateUserTest(user map[string]string, username string, t *testing.T) (status int) {
	db := testInitMigration(t)
	req, err := http.NewRequest("PUT", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockUpdateUser(w, r, user, username, db, t)
	})

	handler.ServeHTTP(rr, req)

	return rr.Code
}

func deleteUserTest(username string, t *testing.T) (status int) {
	db := testInitMigration(t)
	req, err := http.NewRequest("DELETE", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockDeleteUser(w, r, username, db, t)
	})

	handler.ServeHTTP(rr, req)

	return rr.Code
}

func getTopUsersTest(limit string, t *testing.T) (status int, users []u.User) {
	db := testInitMigration(t)
	req, err := http.NewRequest("GET", "/api/leaderboard/{limit}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockGetTopUsers(w, r, limit, db, t)
	})

	handler.ServeHTTP(rr, req)

	status = rr.Code
	json.NewDecoder(rr.Result().Body).Decode(&users)

	return status, users
}
