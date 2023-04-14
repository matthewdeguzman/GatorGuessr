package endpoints

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	u "github.com/matthewdeguzman/GatorGuessr/src/server/structs"
	"gorm.io/gorm"
)

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie from the request using its name.
	// If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	cookieName := ""
	json.NewDecoder(r.Body).Decode(&cookieName)
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie \""+cookieName+"\" not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte(cookie.Value))
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Initialize a new cookie where the name is based on the user ID

	var user u.User
	DecodeUser(&user, r)
	FetchUser(db, &user, user.Username)
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

	w.Write([]byte("Cookie created"))
}
