/**
* Referenced from https://www.alexedwards.net/blog/working-with-cookies-in-go
 */
package cookies

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strconv"

	u "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
)

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

func WriteSignedCookie(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	// Calculate the HMAC signature of the cookie name and value, using SHA256 and
	// a secret key
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)

	// Prepend the cookie value with the HMAC signature.
	cookie.Value = string(signature) + cookie.Value
	log.Println("Written value: " + cookie.Value)

	// encode the value in base64
	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

	// return error if the cookie is too long
	if len(cookie.String()) > 4096 {
		return ErrValueTooLong
	}

	http.SetCookie(w, &cookie)
	w.Write([]byte(cookie.Value))
	return nil
}

func ReadSignedCookie(r *http.Request, name string, secretKey []byte) (string, error) {

	// Read in the signed value from the cookie. This should be in the format
	// "{signature}{original value}".
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}

	// Decode the base64-encoded cookie value. If the cookie didn't contain a
	// valid base64-encoded value, this operation will fail and we return an
	// ErrInvalidValue error.
	// Ensure signedValue is within the proper bound of a sha256 HMAC
	// signature, which has a fixed length of 32 bytes
	signedValue, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil || len(signedValue) < sha256.Size {
		return "", ErrInvalidValue
	}

	signature := signedValue[:sha256.Size]
	value := signedValue[sha256.Size:]

	// Recalculate the HMAC signature of the cookie name and original value.
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(name))
	mac.Write([]byte(value))
	expectedSignature := mac.Sum(nil)

	log.Println("Signature: " + string(signature))
	log.Println("Expected signature: " + string(expectedSignature))
	// If the signatures do not match, then the cookie is not valid
	// and may have been modified by the client
	if !hmac.Equal([]byte(signature), expectedSignature) {
		return "", ErrInvalidValue
	}

	return string(value), nil
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request, user u.User, secretKey []byte) {
	// Initialize a new cookie where the name is based on the user ID

	cookie := http.Cookie{
		Name:   "UserLoginCookie",
		Value:  "UserLogin" + strconv.FormatUint(uint64(user.ID), 10),
		MaxAge: 60 * 60 * 24 * 365 * 5,
		Path:   "/api/",
	}
	err := WriteSignedCookie(w, cookie, secretKey)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request, secretKey []byte) error {
	// Retrieve the cookie from the request using its name.
	// If no matching cookie is found, this will return a
	// http.ErrNoCookie error.
	_, err := ReadSignedCookie(r, "UserLoginCookie", secretKey)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusNotFound)
		case errors.Is(err, ErrInvalidValue):
			http.Error(w, "invalid cookie", http.StatusForbidden)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return err
	}

	return nil
}
