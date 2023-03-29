package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	u "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
	"gorm.io/gorm"
)

func GetUsers(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	SetHeader(w)
	var users []u.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	SetHeader(w)

	params := mux.Vars(r)
	var user u.User
	FetchUser(db, &user, params["username"])

	if user.Username == "" {
		UserDNErr(w)
		return
	}
	EncodeUser(user, w)
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	SetHeader(w)

	var user u.User
	DecodeUser(&user, r)

	if UserExists(db, user.Username) {
		WriteErr(w, http.StatusBadRequest, "400 - User already exists")
		return
	}

	if user.ID != 0 || user.Password == "" {
		WriteErr(w, http.StatusBadRequest, "400 - Attempting to change ID or password is empty")
		return
	}

	hash, err := EncodePassword(user.Password)

	if err != nil {
		WriteErr(w, http.StatusInternalServerError, hashErr)
		return
	}
	user.Password = hash

	db.Create(&user)
	EncodeUser(user, w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	SetHeader(w)
	params := mux.Vars(r)

	var oldUser u.User
	var updatedUser u.User

	FetchUser(db, &oldUser, params["username"])
	FetchUser(db, &updatedUser, params["username"])

	if oldUser.Username == "" {
		UserDNErr(w)
		return
	}

	DecodeUser(&updatedUser, r)

	if oldUser.ID != updatedUser.ID {
		WriteErr(w, http.StatusMethodNotAllowed, "405 - Cannot change immutable field")
		return
	}

	hash, err := EncodePassword(updatedUser.Password)
	if err != nil {
		HashErr(w)
		return
	}
	updatedUser.Password = hash
	updatedUser.CreatedAt = oldUser.CreatedAt

	db.Save(&updatedUser)
	EncodeUser(updatedUser, w)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	SetHeader(w)
	params := mux.Vars(r)
	var user u.User

	FetchUser(db, &user, params["username"])
	if user.Username == "" {
		UserDNErr(w)
		return
	}
	db.Delete(&user, "Username = ?", params["username"])
	EncodeUser(user, w)
}

func ValidateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	SetHeader(w)

	var user u.User
	var givenPassword string
	var hashedPassword string

	DecodeUser(&user, r)
	givenPassword = user.Password
	FetchUser(db, &user, user.Username)

	if user.Username == "" {
		UserDNErr(w)
		return
	}
	hashedPassword = user.Password

	match, err := DecodePasswordAndMatch(givenPassword, hashedPassword)
	if err != nil {
		HashErr(w)
		return
	}
	if !match {
		LoginErr(w)
		return
	}
}
