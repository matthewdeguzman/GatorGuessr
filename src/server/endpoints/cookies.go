package endpoints

import (
	"errors"
	"log"
	"net/http"
)

const cookieName = "user-cookie"

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize a new cookie containing the string "Hello world!" and some
	// non-default attributes.
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    "",
		MaxAge:   60 * 60 * 24 * 365 * 5, // 5 years
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	http.SetCookie(w, &cookie)

	// Write a HTTP response as normal.
	w.Write([]byte("cookie set"))
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie from the request using its name.
	// If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	// Echo out the cookie value in the response body.
	w.Write([]byte(cookie.Value))
}
