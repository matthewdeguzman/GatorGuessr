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

	cleanDB(&user, user.Username)
}

func TestCreateUserWithoutPassword(t *testing.T) {
	user := u.User{
		Username: "newuseralert!!",
	}

	if status := createUserTest(user, t); status != http.StatusBadRequest {
		t.Fail()
	}

	cleanDB(&user, user.Username)
}
func TestCreateWithID(t *testing.T) {
	t.Fail()
}

func TestUpdateNonexistantUser(t *testing.T) {
	t.Fail()
}

func TestUpdateExistingUser(t *testing.T) {
	t.Fail()
}

func TestUpdateUserID(t *testing.T) {
	t.Fail()
}

func TestUpdateUserScore(t *testing.T) {
	t.Fail()
}

func TestDeleteExistingUser(t *testing.T) {
	t.Fail()
}

func TestNonexistingUser(t *testing.T) {
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
