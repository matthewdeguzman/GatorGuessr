package endpoints

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	u "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
	"gorm.io/gorm"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Initialize a new cookie containing the string "Hello world!" and some
	// non-default attributes.

	params := mux.Vars(r)
	var user u.User
	FetchUser(db, &user, params["username"])
	if user.Username == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	cookieName := "cookie" + strconv.FormatUint(uint64(user.ID), 10)
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    "user cookie",
		MaxAge:   60 * 60 * 24 * 365 * 5, // 5 years
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	http.SetCookie(w, &cookie)

	// Write a HTTP response as normal.
	w.Write([]byte("cookie name: " + cookieName))
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie from the request using its name.
	// If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	params := mux.Vars(r)
	cookie, err := r.Cookie(params["cookie_name"])
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie \""+params["cookie_name"]+"\" not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	// Echo out the cookie value in the response body.
	w.Write([]byte(cookie.Value))
}
