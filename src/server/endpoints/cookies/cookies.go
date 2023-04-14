/**
* Referenced from https://www.alexedwards.net/blog/working-with-cookies-in-go
 */
package cookies

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"
)

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

func writeCookie(w http.ResponseWriter, cookie http.Cookie) error {

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

func readCookie(r *http.Request, name string) (string, error) {

	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}

	// Decode the base64-encoded cookie value. If the cookie didn't contain a
	// valid base64-encoded value, this operation will fail and we return an
	// ErrInvalidValue error.
	value, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return "", ErrInvalidValue
	}

	// Return the decoded cookie value.
	return string(value), nil
}

func WriteSignedCookie(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	// Calculate the HMAC signature of the cookie name and value, using SHA256 and
	// a secret key
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)

	// Prepend the cookie value with the HMAC signature.
	cookie.Value = string(signature) + cookie.Value

	return writeCookie(w, cookie)
}

func ReadSignedCookie(r *http.Request, name string, secretKey []byte) (string, error) {

	// Read in the signed value from the cookie. This should be in the format
	// "{signature}{original value}".
	signedValue, err := readCookie(r, name)
	if err != nil {
		return "", err
	}

	// Ensure signedValue is within the proper bound of a sha256 HMAC
	// signature, which has a fixed length of 32 bytes
	if len(signedValue) < sha256.Size {
		return "", ErrInvalidValue
	}

	signature := signedValue[:sha256.Size]
	value := signedValue[sha256.Size:]

	// Recalculate the HMAC signature of the cookie name and original value.
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(name))
	mac.Write([]byte(value))
	expectedSignature := mac.Sum(nil)

	// If the signatures do not match, then the cookie is not valid
	// and may have been modified by the client
	if !hmac.Equal([]byte(signature), expectedSignature) {
		return "", ErrInvalidValue
	}

	return value, nil
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, secretKey []byte) {
	// Initialize a new cookie where the name is based on the user ID
	var cookie http.Cookie
	json.NewDecoder(r.Body).Decode(&cookie)
	err := WriteSignedCookie(w, cookie, secretKey)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Cookie created"))
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request, secretKey []byte) {
	// Retrieve the cookie from the request using its name.
	// If no matching cookie is found, this will return a
	// http.ErrNoCookie error.

	var cookie http.Cookie
	json.NewDecoder(r.Body).Decode(&cookie)
	value, err := ReadSignedCookie(r, cookie.Name, secretKey)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		case errors.Is(err, ErrInvalidValue):
			http.Error(w, "invalid cookie", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte(value))
}
