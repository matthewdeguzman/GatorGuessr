package endpoints

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	u "github.com/matthewdeguzman/GatorGuessr/src/server/db_user"
	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

var (
	HashErr   = "500 - Error with hashing. User not created."
	UserDNErr = "404 - User not found."
	LoginErr  = "404 - Username or Password Incorrect."
)

type hashParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func writeErr(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func hashErr(w http.ResponseWriter) {
	writeErr(w, http.StatusInternalServerError, HashErr)
}

func userDNErr(w http.ResponseWriter) {
	writeErr(w, http.StatusNotFound, UserDNErr)
}

func loginErr(w http.ResponseWriter) {
	writeErr(w, http.StatusNotFound, LoginErr)
}

func userExists(db *gorm.DB, username string) bool {
	var user u.User
	db.First(&user, "Username = ?", username)
	if user.Username == "" {
		return false
	} else {
		return true
	}
}

func fetchUser(db *gorm.DB, user *u.User, username string) {
	db.First(user, "Username = ?", username)
}

func decodeUser(user *u.User, r *http.Request) {
	json.NewDecoder(r.Body).Decode(user)
}

func encodeUser(user u.User, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(user)
}

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
