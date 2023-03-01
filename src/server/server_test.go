package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func testGetPassword(t *testing.T) string {
	// Get password from credentials.txt
	if err != nil {
		panic(err)
	}
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(filename)))), "credentials.txt")
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	password := ""
	for scanner.Scan() {
		password += scanner.Text()
	}
	password = strings.TrimSpace(password)

	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	return password
}

func testInitMigration(t *testing.T) {
	const DB_USERNAME = "cen3031"
	const DB_NAME = "user_database"
	const DB_HOST = "cen3031-server.mysql.database.azure.com"
	const DB_PORT = "3306"
	var password = testGetPassword(t)
	// Build connection string
	DSN := DB_USERNAME + ":" + password + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	// migrates the server if necessary
	db.AutoMigrate(&User{})
}

func getUserTest(username string, t *testing.T) string {
	testInitMigration(t)
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Error(err)
	}

	mockGetUser := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := username
		var user User
		db.First(&user, "Username = ?", params)
		json.NewEncoder(w).Encode(user)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockGetUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	var user User
	if err := json.Unmarshal(rr.Body.Bytes(), &user); err != nil {
		t.Errorf("got invalid reponse, expected a user, got: %v", rr.Body.String())
	}

	return user.Username
}

func deleteUserTest(username string, t *testing.T) {
	testInitMigration(t)
	req, err := http.NewRequest("DELETE", "/api/users", nil)
	if err != nil {
		t.Error(err)
	}

	mockDeleteUser := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := username
		var user User
		db.First(&user, "Username = ?", params)
		db.Delete(&user)
		json.NewEncoder(w).Encode(user)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockDeleteUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	var user User
	if err := json.Unmarshal(rr.Body.Bytes(), &user); err != nil {
		t.Errorf("got invalid reponse, expected a user, got: %v", rr.Body.String())
	}
}

func createTest(username string, t *testing.T) {
	testInitMigration(t)
	req, err := http.NewRequest("POST", "/api/users", nil)
	if err != nil {
		t.Error(err)
	}

	mockCreateUser := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user User
		user.Username = username
		user.Password = "test-password"
		db.Create(&user)
		json.NewEncoder(w).Encode(user)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockCreateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	var user User
	if err := json.Unmarshal(rr.Body.Bytes(), &user); err != nil {
		t.Errorf("got invalid reponse, expected a user, got: %v", rr.Body.String())
	}
}

func updateUserTest(username string, password string, t *testing.T) {
	testInitMigration(t)
	req, err := http.NewRequest("PUT", "/api/users/{username}", nil)
	if err != nil {
		t.Error(err)
	}

	mockUpdateUser := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user User
		db.First(&user, "Username = ?", username)
		user.Password = password
		db.Save(&user)
		json.NewEncoder(w).Encode(user)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockUpdateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	var user User
	if err := json.Unmarshal(rr.Body.Bytes(), &user); err != nil {
		t.Errorf("got invalid reponse, expected a user, got: %v", rr.Body.String())
	}
}

/// TESTS ///

func TestGetUsers(t *testing.T) {

	// initializes the db and sends the get request
	testInitMigration(t)
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Error(err)
	}

	// creates rr to get the response recorder and makes the handler for the getUser api
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsers)

	// passes in the response recorder and the request
	handler.ServeHTTP(rr, req)

	// if the status code is not expected, we error
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	// unmarshal json data into users array
	var users []User
	if err := json.Unmarshal(rr.Body.Bytes(), &users); err != nil {
		t.Errorf("got invalid response, expected list of users, got: %v", rr.Body.String())
	}
}

func TestGetUser1(t *testing.T) {
	username := "matthew"
	if resp := getUserTest(username, t); resp != username {
		t.Errorf("got invalid response, expected %v, got: %v", username, resp)
	}
}

func TestGetUser2(t *testing.T) {
	username := "invalid-user"
	if resp := getUserTest(username, t); resp != "" {
		t.Errorf("got invalid response, expected %v, got: %v", username, resp)
	}
}

func TestCreateUser(t *testing.T) {
	createTest("test-user", t)
}

func TestDeleteUser1(t *testing.T) {
	if resp := getUserTest("test-user", t); resp == "test-user" {
		deleteUserTest("test-user", t)
	} else {
		createTest("test-user", t)
		deleteUserTest("test-user", t)
	}
}

func TestDeleteUser2(t *testing.T) {
	if resp := getUserTest("test-user", t); resp == "test-user" {
		deleteUserTest("test-user", t)
		deleteUserTest("test-user", t)
	} else {
		deleteUserTest("test-user", t)
	}
}

func TestUpdateUser1(t *testing.T) {
	if resp := getUserTest("test-user", t); resp == "test-user" {
		deleteUserTest("test-user", t)
	}
	createTest("test-user", t)
	updateUserTest("test-user", "new-password", t)
	deleteUserTest("test-user", t)

}

func TestUpdateUser2(t *testing.T) {
	deleteUserTest("test-user", t)
	updateUserTest("test-user", "new-password", t)
}
