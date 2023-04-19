package tests

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/matthewdeguzman/GatorGuessr/src/server/endpoints"
	"github.com/matthewdeguzman/GatorGuessr/src/server/endpoints/api"
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
	cleanDB(user, db)
	db.Create(&user)
	return nil
}

func cookieExists(name string, cookies []*http.Cookie, t *testing.T) bool {
	// if there is no cookie with the expected name, then the test fails
	for _, cookie := range cookies {
		if cookie.Name == "UserLoginCookie" {
			return true
		}
	}

	return false

}

func addCookie(user u.User, r *http.Request, secretKey []byte) {

	cookie := http.Cookie{
		Name:   "UserLoginCookie",
		Value:  "UserLogin" + strconv.FormatUint(uint64(user.ID), 10),
		MaxAge: 60 * 60 * 24 * 365 * 5,
		Path:   "/api/",
	}

	// Calculate the HMAC signature of the cookie name and value, using SHA256 and
	// a secret key
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)

	// Prepend the cookie value with the HMAC signature.
	cookie.Value = string(signature) + cookie.Value
	log.Println("Written value: " + cookie.Value)

	// encode the value in base64
	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

	r.AddCookie(&cookie)

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

func getUserTest(user u.User, t *testing.T, db *gorm.DB, secretKey []byte) (status int) {
	req, err := http.NewRequest("GET", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addCookie(user, r, secretKey)
		api.GetUserWithUsername(w, r, user.Username, db, secretKey)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
}

func createUserTest(user u.User, t *testing.T, db *gorm.DB, secretKey []byte) (status int) {

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
		api.CreateUser(w, r, db, secretKey)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
}

func updateUserTest(ogUser, updatedUser u.User, t *testing.T, db *gorm.DB, secretKey []byte) (status int) {

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
		addCookie(ogUser, r, secretKey)
		api.UpdateUserFromUser(w, r, ogUser, db, secretKey)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
}

func deleteUserTest(user u.User, t *testing.T, db *gorm.DB, secretKey []byte) (status int) {
	req, err := http.NewRequest("DELETE", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addCookie(user, r, secretKey)
		api.DeleteUserFromUsername(w, r, user, db, secretKey)
	})

	handler.ServeHTTP(rr, req)

	return rr.Result().StatusCode
}

func validateUserTest(sentUser u.User, t *testing.T, db *gorm.DB, secretKey []byte) (status int) {

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
		addCookie(sentUser, r, secretKey)
		api.ValidateUser(w, r, db, secretKey)
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
