package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {

	// initializes the db and sends the get request
	initializeMigration()
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

func userTest(username string, t *testing.T) {
	initializeMigration()
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

	if user.Username != username {
		t.Errorf("got invalid user, expected %v, got: %v", username, user.Username)
	}
}
func TestGetUser1(t *testing.T) {
	userTest("matthew", t)
}

func TestGetUser2(t *testing.T) {
	userTest("ethanfan", t)
}

func TestGetUser3(t *testing.T) {
	userTest("stephencoomes", t)
}

func TestGetUser4(t *testing.T) {
	userTest("matthew", t)
}
