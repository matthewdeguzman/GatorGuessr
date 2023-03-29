package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	u "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
)

/// TESTS ///

func TestGetUsers(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/users/", nil)
	if err != nil {
		t.Error(err)
	}

	// creates rr to get the response recorder and makes the handler for the getUser api
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockGetUsers(w, r, t)
	})

	// passes in the response recorder and the request
	handler.ServeHTTP(rr, req)

	// if the status code is not expected, we error
	if status := rr.Code; status != http.StatusOK {
		t.Error(string(rune(status)))
	}
}

func TestGetExistingUser(t *testing.T) {
	if status := getUserTest("matthew", t); status != http.StatusOK {
		t.Fail()
	}
}

func TestGetNonexiststantUser(t *testing.T) {
	if status := getUserTest("jksdal;fjea;ils", t); status != http.StatusNotFound {
		t.Fail()
	}
}

func TestCreateExistingUser(t *testing.T) {
	user := u.User{
		Username: "matthew",
		Password: "passwordddd",
	}
	if status := createUserTest(user, t); status != http.StatusBadRequest {
		cleanDB(&user, user.Username, t)
		t.Log(status)
		t.Fail()
	}
}

func TestCreateNewUser(t *testing.T) {
	user := u.User{
		Username: "garbledmess",
		Password: "password!!",
	}

	if status := createUserTest(user, t); status != http.StatusOK {
		t.Fail()
	}
	cleanDB(&user, user.Username, t)
}

func TestCreateUserWithoutPassword(t *testing.T) {
	user := u.User{
		Username: "newuseralert!!",
	}

	if status := createUserTest(user, t); status != http.StatusBadRequest {
		t.Fail()
		cleanDB(&user, user.Username, t)
	}

}
func TestCreateUserWithID(t *testing.T) {
	user := u.User{
		Username: "newuseralert!!",
		Password: "pasworddd",
		ID:       10,
	}

	if status := createUserTest(user, t); status != http.StatusBadRequest {
		t.Fail()
		cleanDB(&user, user.Username, t)
	}

}

func TestUpdateNonexistantUser(t *testing.T) {
	user := map[string]string{
		"Username": "this user doesn't exist lol",
		"Password": "yeah",
	}

	status := updateUserTest(user, user["Username"], t)
	if status != http.StatusNotFound {
		t.Log(status)
		t.Fail()
	}
	cleanDB(&u.User{}, user["Username"], t)
}

func TestUpdateExistingUser(t *testing.T) {
	user := map[string]string{
		"Username": "matthew",
		"Password": "yeah",
	}

	status := updateUserTest(user, user["Username"], t)
	if status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}
}

func TestUpdateUserID(t *testing.T) {
	user := map[string]string{
		"ID":       "88349",
		"Username": "matthew",
	}

	status := updateUserTest(user, user["Username"], t)
	if status != http.StatusMethodNotAllowed {
		t.Log(status)
		t.Fail()
	}
}

func TestUpdateUserScore(t *testing.T) {
	user := map[string]string{
		"Username": "stephen",
		"Score":    "0",
	}

	status := updateUserTest(user, user["Username"], t)
	if status != http.StatusOK {
		t.Log(status)
		t.Fail()
	}
}

func TestDeleteExistingUser(t *testing.T) {
	t.Fail()
}

func TestDeleteNonexistingUser(t *testing.T) {
	t.Fail()
}

func TestValidateExistingUser(t *testing.T) {
	t.Fail()
}

func TestValidNonexistantuser(t *testing.T) {
	t.Fail()
}

func TestLeaderboardNegativeInteger(t *testing.T) {
	t.Fail()
}

func TestLeaderboardFloatintPoint(t *testing.T) {
	t.Fail()
}

func TestLeaderboardZero(t *testing.T) {
	t.Fail()
}

func TestLeaderboardOverMaxLimit(t *testing.T) {
	t.Fail()
}

func TestLeaderboardCasual(t *testing.T) {
	t.Fail()
}
func TestPasswordEncoding(t *testing.T) {
	t.Fail()
}

func TestPasswordDecoding(t *testing.T) {
	t.Fail()
}
