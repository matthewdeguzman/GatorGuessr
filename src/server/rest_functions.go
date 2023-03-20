package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"crypto/rand"
	"crypto/subtle"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/argon2"
)

type hashParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var (
	HashErr   = []byte("500 - Error with hashing. User not created.")
	UserDNErr = []byte("404 - User not found.")
	LoginErr  = []byte("404 - Username or Password Incorrect.")
)

func hashErr(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(HashErr)
}

func userDNErr(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write(UserDNErr)
}

func loginErr(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write(LoginErr)
}

// generates a hashed version of the given password using argon2
func encodePassword(password string) (encodedHash string, err error) {

	p := &hashParams{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	salt := make([]byte, p.saltLength)
	_, err = rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func decodePasswordAndMatch(password, encodedHash string) (match bool, err error) {
	var (
		ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
		ErrIncompatibleVersion = errors.New("incompatible version of argon2")
	)

	// get hash
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return false, ErrInvalidHash
	}

	// get version
	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return false, err
	}
	if version != argon2.Version {
		return false, ErrIncompatibleVersion
	}

	// get hash parameters
	p := &hashParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return false, err
	}

	// get salt
	salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return false, err
	}
	p.saltLength = uint32(len(salt))

	// get hash
	hash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return false, err
	}
	p.keyLength = uint32(len(hash))

	// hash password
	hashPass := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// compare hashPass to the given hash. we use subtle bc it helps
	// prevents against timing attacks
	if subtle.ConstantTimeCompare(hash, hashPass) == 1 {
		return true, nil
	}

	// the passwords do not match
	return false, nil
}

func EnableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

// GetUsers returns all the users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// GetUser returns a specified user from the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var user User
	db.First(&user, "Username = ?", params["username"])

	// if user does not exist, the username will be empty, so
	// we send back an invalid request
	if user.Username == "" {
		userDNErr(w)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// createUser creates a new user and inserts into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// decodes user
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	// use argon2 to hash the passwords
	hash, err := encodePassword(user.Password)

	// if there is an error, respond with an internal server error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(HashErr)
		return
	}
	user.Password = hash // stores password as encoded password

	// create and encode the user
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// updateUser updates a user with the sent information
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // receives the given parameters in the request

	// finds the user based on the parameters
	var user User
	db.First(&user, "Username = ?", params["username"])

	// if the username is empty, then the user does not exist
	// respond with a 400 bad request error
	if user.Username == "" {
		userDNErr(w)
		return
	}

	// decode the user
	json.NewDecoder(r.Body).Decode(&user)

	// use argon2 to hash the passwords
	hash, err := encodePassword(user.Password)

	// if there is an error with hashing, respond with error
	if err != nil {
		hashErr(w)
		return
	}

	// store hashed password then save and encode the user
	user.Password = hash
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// deleteUser deletes a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.Delete(&user, "Username = ?", params["username"])
	json.NewEncoder(w).Encode(user)
}

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// decode user object and store the given password
	var user User
	var givenPassword string
	var hashedPassword string

	json.NewDecoder(r.Body).Decode(&user)
	givenPassword = user.Password

	// retrieve user from db
	db.First(&user, "Username = ?", user.Username)

	// if user does not exist, the username will be empty, so
	// we send back an invalid request
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Username and Password Match"))
}
