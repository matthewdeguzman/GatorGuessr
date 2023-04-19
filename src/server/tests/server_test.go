package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthewdeguzman/GatorGuessr/src/server/endpoints/api"
	u "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
)

/// TESTS ///

func TestGetUsers(t *testing.T) {
	db := testInitMigration(t)
	req, err := http.NewRequest("GET", "/api/users/", nil)
	if err != nil {
		t.Error(err)
	}

	// creates rr to get the response recorder and makes the handler for the getUser api
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.GetUsers(w, r, db)
	})

	// passes in the response recorder and the request
	handler.ServeHTTP(rr, req)

	// if the status code is not expected, we error
	if status := rr.Code; status != http.StatusOK {
		t.Error(string(rune(status)))
	}
}

func TestWithoutAuthorization(t *testing.T) {

	db := testInitMigration(t)
	user := u.User{
		Username: "User",
		Password: "User",
	}

	addUser(user, t, db)

	req, err := http.NewRequest("GET", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	// sends request without setting cookie, should return a forbidden status
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.GetUserWithUsername(w, r, user.Username, db)
	})

	handler.ServeHTTP(rr, req)

	status := rr.Result().StatusCode

	if status != http.StatusForbidden {
		t.Fail()
	}
}
func TestGetExistingUser(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "User",
		Password: "User",
	}

	addUser(user, t, db)
	status := getUserTest(user, t, db)
	cleanDB(user, db)

	if status != http.StatusOK {
		t.Fail()
	}
}

func TestGetNonexiststantUser(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "NonexistantUser",
		Password: "Password",
	}

	cleanDB(user, db)
	status := getUserTest(user, t, db)

	if status != http.StatusNotFound {
		t.Fail()
	}
}

func TestCreateExistingUser(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "User",
		Password: "User",
	}

	addUser(user, t, db)
	status := createUserTest(user, t, db)
	cleanDB(user, db)

	if status != http.StatusBadRequest {
		t.Log(status)
		t.Fail()
	}
}

func TestCreateNewUser(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "NewUser",
		Password: "User",
	}
	cleanDB(user, db)
	if status := createUserTest(user, t, db); status != http.StatusOK {
		t.Fail()
	}
	cleanDB(user, db)
}

func TestCreateUserWithoutPassword(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "NewUser",
	}
	cleanDB(user, db)
	if status := createUserTest(user, t, db); status != http.StatusBadRequest {
		t.Fail()
	}

}
func TestCreateUserWithID(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "NewUser",
		Password: "Password",
		ID:       10,
	}
	cleanDB(user, db)
	if status := createUserTest(user, t, db); status != http.StatusBadRequest {
		t.Fail()
	}

}

func TestUpdateUserWithoutAuthorization(t *testing.T) {
	db := testInitMigration(t)
	ogUser := u.User{
		Username: "User",
		Password: "User",
	}
	updatedUser := u.User{
		Password: "NewPassword",
	}
	cleanDB(ogUser, db)
	addUser(ogUser, t, db)

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

	// attempts to update user without cookie. should return forbidden request status
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.UpdateUserFromUser(w, r, ogUser, db)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusForbidden {
		t.Log(status)
		t.Fail()
	}
}
func TestUpdateNonexistantUser(t *testing.T) {
	db := testInitMigration(t)
	ogUser := u.User{
		Username: "NonexistantUser",
		Password: "password",
	}

	cleanDB(ogUser, db)
	status := updateUserTest(u.User{}, ogUser, t, db)
	if status != http.StatusNotFound {
		t.Log(status)
		t.Fail()
	}
}

func TestUpdateExistingUser(t *testing.T) {
	db := testInitMigration(t)
	ogUser := u.User{
		Username: "User",
		Password: "User",
	}
	updatedUser := u.User{
		Password: "NewPassword",
	}
	cleanDB(ogUser, db)
	addUser(ogUser, t, db)
	status := updateUserTest(ogUser, updatedUser, t, db)
	if status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}
}

func TestUpdateUserID(t *testing.T) {
	db := testInitMigration(t)
	ogUser := u.User{
		Username: "User",
		Password: "User",
	}
	updatedUser := u.User{
		ID: 9403059,
	}

	cleanDB(ogUser, db)
	addUser(ogUser, t, db)
	status := updateUserTest(ogUser, updatedUser, t, db)

	if status != http.StatusMethodNotAllowed {
		t.Log(status)
		t.Fail()
	}
}

func TestDeleteUserWithoutAuthorization(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "User",
		Password: "User",
	}

	cleanDB(user, db)
	addUser(user, t, db)
	req, err := http.NewRequest("DELETE", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.DeleteUserFromUsername(w, r, user, db)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}
}

func TestDeleteExistingUser(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "User",
		Password: "User",
	}

	cleanDB(user, db)
	addUser(user, t, db)
	status := deleteUserTest(user, t, db)
	if status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}

}

func TestValidateExistantUserCookie(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "User",
		Password: "User",
	}
	cleanDB(user, db)
	addUser(user, t, db)

	sentMarshal, err := json.Marshal(user)
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

	// if there is no cookie with the expected name, then the test fails
	var cookieExists bool = false
	for _, cookie := range rr.Result().Cookies() {
		if cookie.Name == "UserLoginCookie" {
			cookieExists = true
			break
		}
	}
	if !cookieExists {
		t.Log("Cookie not set")
		t.Fail()
	}

	if status := rr.Result().StatusCode; status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}
}
func TestValidateExistingUser(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "User",
		Password: "User",
	}
	cleanDB(user, db)
	addUser(user, t, db)

	status := validateUserTest(user, t, db)
	if status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}
}

func TestValidateNonexistantuser(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "User",
		Password: "User",
	}
	cleanDB(user, db)
	status := validateUserTest(user, t, db)
	if status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}
}

func TestValidateIncorrectPassword(t *testing.T) {
	db := testInitMigration(t)
	realUser := u.User{
		Username: "User",
		Password: "User",
	}
	sentUser := u.User{
		Username: realUser.Username,
		Password: "WrongPassword",
	}

	cleanDB(realUser, db)
	addUser(realUser, t, db)

	status := validateUserTest(sentUser, t, db)

	if status != http.StatusNotFound {
		t.Log(status)
		t.Fail()
	}
}

func TestLeaderboardNegativeInteger(t *testing.T) {
	limit := "-248"
	status, _ := getTopUsersTest(limit, t)

	if status != http.StatusBadRequest {
		t.Log(status)
		t.Fail()
	}
}

func TestLeaderboardFloatintPoint(t *testing.T) {
	limit := "3.534"
	status, _ := getTopUsersTest(limit, t)

	if status != http.StatusBadRequest {
		t.Log(status)
		t.Fail()
	}
}

func TestLeaderboardZero(t *testing.T) {
	limit := "0"
	status, _ := getTopUsersTest(limit, t)

	if status != http.StatusBadRequest {
		t.Log(status)
		t.Fail()
	}
}

func TestLeaderboardOverMaxLimit(t *testing.T) {
	limit := "99999"
	status, _ := getTopUsersTest(limit, t)

	if status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}
}

func TestLeaderboardSorted(t *testing.T) {
	limit := "10"
	status, users := getTopUsersTest(limit, t)

	if status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}

	var previousScore uint = users[0].Score
	for _, user := range users {
		if user.Score > previousScore {
			t.Log("Not sorted in descending order")
			t.Fail()
		}
	}
}
