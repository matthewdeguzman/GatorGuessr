package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/matthewdeguzman/GatorGuessr/src/server/endpoints"
	"github.com/matthewdeguzman/GatorGuessr/src/server/endpoints/api"
	"github.com/matthewdeguzman/GatorGuessr/src/server/endpoints/cookies"
	u "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func testInitMigration(t *testing.T) (db *gorm.DB) {
	const DB_USERNAME = "cen3031"
	const DB_NAME = "test_database"
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

func cleanDB(user u.User, db *gorm.DB) {
	db.Delete(user, "Username = ?", user.Username)
}

func addUser(user u.User, t *testing.T, db *gorm.DB) (err error) {
	hash, err := endpoints.EncodePassword(user.Password)

	if err != nil {
		t.Error(err)
	}
	user.Password = hash

	db.Create(&user)
	return nil
}

func cookieExists(name string, cookies []*http.Cookie) bool {
	// if there is no cookie with the expected name, then the test fails
	for _, cookie := range cookies {
		if cookie.Name == "UserLoginCookie" {
			return true
		}
	}

	return false

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

func getUserTest(user u.User, t *testing.T, db *gorm.DB) (status int) {
	req, err := http.NewRequest("GET", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies.SetCookieHandler(w, r, user)
		api.GetUserWithUsername(w, r, user.Username, db)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
}

func createUserTest(user u.User, t *testing.T, db *gorm.DB) (status int) {

	marshal, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("POST", "/api/users/", bytes.NewReader(marshal))
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.CreateUser(w, r, db)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
}

func updateUserTest(ogUser, updatedUser u.User, t *testing.T, db *gorm.DB) (status int) {

	updatedMarshal, err := json.Marshal(updatedUser)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("PUT", "/api/users/{username}/", bytes.NewReader(updatedMarshal))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies.SetCookieHandler(w, r, ogUser)
		api.UpdateUserFromUser(w, r, ogUser, db)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
}

func deleteUserTest(user u.User, t *testing.T, db *gorm.DB) (status int) {
	req, err := http.NewRequest("DELETE", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies.SetCookieHandler(w, r, user)
		api.DeleteUserFromUsername(w, r, user, db)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
}

func validateUserTest(sentUser u.User, t *testing.T, db *gorm.DB) (status int) {

	sentMarshal, err := json.Marshal(sentUser)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("POST", "/api/login/", bytes.NewReader(sentMarshal))
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.ValidateUser(w, r, db)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
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
