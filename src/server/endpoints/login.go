package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	u "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
	"gorm.io/gorm"
)

func GetUsers(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	setHeader(w)
	var users []u.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	setHeader(w)

	params := mux.Vars(r)
	var user u.User
	fetchUser(db, &user, params["username"])

	if user.Username == "" {
		userDNErr(w)
		return
	}
	encodeUser(user, w)
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	setHeader(w)

	var user u.User
	decodeUser(&user, r)

	if userExists(db, user.Username) {
		writeErr(w, http.StatusBadRequest, "400 - User already exists")
		return
	}

	hash, err := encodePassword(user.Password)

	if err != nil {
		writeErr(w, http.StatusInternalServerError, HashErr)
		return
	}
	user.Password = hash

	db.Create(&user)
	encodeUser(user, w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	setHeader(w)
	params := mux.Vars(r)

	var oldUser u.User
	var updatedUser u.User

	fetchUser(db, &oldUser, params["username"])
	fetchUser(db, &updatedUser, params["username"])

	if oldUser.Username == "" {
		userDNErr(w)
		return
	}

	decodeUser(&updatedUser, r)

	if oldUser.ID != updatedUser.ID {
		writeErr(w, http.StatusMethodNotAllowed, "405 - Cannot change immutable field")
		return
	}

	hash, err := encodePassword(updatedUser.Password)
	if err != nil {
		hashErr(w)
		return
	}
	updatedUser.Password = hash
	updatedUser.CreatedAt = oldUser.CreatedAt

	db.Save(&updatedUser)
	encodeUser(updatedUser, w)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	setHeader(w)
	params := mux.Vars(r)
	var user u.User

	fetchUser(db, &user, params["username"])
	if user.Username == "" {
		userDNErr(w)
		return
	}
	db.Delete(&user, "Username = ?", params["username"])
	encodeUser(user, w)
}

func ValidateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	setHeader(w)

	var user u.User
	var givenPassword string
	var hashedPassword string

	decodeUser(&user, r)
	givenPassword = user.Password
	fetchUser(db, &user, user.Username)

	if user.Username == "" {
		userDNErr(w)
		return
	}
	hashedPassword = user.Password

	match, err := decodePasswordAndMatch(givenPassword, hashedPassword)
	if err != nil {
		hashErr(w)
		return
	}
	if !match {
		loginErr(w)
		return
	}
}
