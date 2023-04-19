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

var secretKey []byte = []byte("test_secret")

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

func TestGetWithoutAuthorization(t *testing.T) {

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

	// sends request without setting cookie, should return a not found status
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.GetUserWithUsername(w, r, user.Username, db, secretKey)
	})

	handler.ServeHTTP(rr, req)

	status := rr.Result().StatusCode

	if status != http.StatusNotFound {
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
	status := getUserTest(user, t, db, secretKey)

	if status != http.StatusOK {
		t.Fail()
	}
}

func TestGetNonexistantUser(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "NonexistantUser",
		Password: "Password",
	}

	cleanDB(user, db)
	status := getUserTest(user, t, db, secretKey)

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
	status := createUserTest(user, t, db, secretKey)
	cleanDB(user, db)

	if status != http.StatusBadRequest {
		t.Log(status)
		t.Fail()
	}
}

func TestCreateNewUser(t *testing.T) {
	db := testInitMigration(t)
	user := u.User{
		Username: "User",
		Password: "User",
	}
	cleanDB(user, db)

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
		if w.Header().Get("Set-Cookie") == "" {
			t.Log("Cookie not properly set")
			t.Fail()
		}
	})

	handler.ServeHTTP(rr, req)
	if status := rr.Result().StatusCode; status != http.StatusOK {
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
	if status := createUserTest(user, t, db, secretKey); status != http.StatusBadRequest {
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
	if status := createUserTest(user, t, db, secretKey); status != http.StatusBadRequest {
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
		api.UpdateUserFromUser(w, r, ogUser, db, secretKey)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusNotFound {
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
	status := updateUserTest(u.User{}, ogUser, t, db, secretKey)
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
	addUser(ogUser, t, db)
	status := updateUserTest(ogUser, updatedUser, t, db, secretKey)
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

	addUser(ogUser, t, db)
	status := updateUserTest(ogUser, updatedUser, t, db, secretKey)

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

	addUser(user, t, db)
	req, err := http.NewRequest("DELETE", "/api/users/{username}/", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.DeleteUserFromUsername(w, r, user, db, secretKey)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusNotFound {
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

	addUser(user, t, db)
	status := deleteUserTest(user, t, db, secretKey)
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
		api.ValidateUser(w, r, db, secretKey)
	})

	handler.ServeHTTP(rr, req)

	// if there is no cookie with the expected name, then the test fails
	if !cookieExists("UserLoginCookie", rr.Result().Cookies(), t) {
		t.Log("Cookie not properly set")
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

	addUser(user, t, db)

	status := validateUserTest(user, t, db, secretKey)
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
	status := validateUserTest(user, t, db, secretKey)
	if status == http.StatusOK {
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

	addUser(realUser, t, db)

	status := validateUserTest(sentUser, t, db, secretKey)

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
	db := testInitMigration(t)
	limit := "10"

	user1 := u.User{
		Username: "User1",
		Password: "password",
		Score:    999,
	}
	user2 := u.User{
		Username: "User1",
		Password: "password",
		Score:    100,
	}
	user3 := u.User{
		Username: "User1",
		Password: "password",
		Score:    5,
	}
	addUser(user1, t, db)
	addUser(user2, t, db)
	addUser(user3, t, db)
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
