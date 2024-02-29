package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
	"log"
	"net/http"
)

const EVENT_URL = "/"
const QUEUEIT_EVENT_ID = "123"

var KEY, _ = ioutil.ReadFile("queueit.pub")

// Middleware function to intercept and process requests
func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the token is present in the query string
		token := fetchToken(r)

		if token != "" && verifyToken(token) {
			setCookie(w, "queueit_token", token)

			// As long as it's a redirect from queueit server, we do a redirect to the server
			if r.URL.Query().Has("queueit_token") {
				// The home url of the event
				http.Redirect(w, r, EVENT_URL, http.StatusFound)
			}
		} else {
			// Redirect the user to http://127.0.0.1:8081 (queueit server)
			http.Redirect(w, r, fmt.Sprintf("http://127.0.0.1:8081/?id=%s", QUEUEIT_EVENT_ID), http.StatusFound)
			return
		}

		// Call the next handler in the chain
		next(w, r)
	}
}

func fetchToken(r *http.Request) string {
	// Check if the token is present in the query string
	token := r.URL.Query().Get("queueit_token")

	// If the token is not found in the query string, check the cookie
	if token == "" {
		cookie, err := r.Cookie("queueit_token")
		if err == nil {
			token = cookie.Value
		}
	}
	return token
}

func setCookie(w http.ResponseWriter, name string, value string) {
	// Create a new cookie
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   1800, // Set the cookie expiration time to 30 minutes in seconds
	}

	// Set the cookie in the response
	http.SetCookie(w, cookie)
}

func verifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(KEY)
	})

	if err != nil {
		log.Printf(err.Error())
		return false
	}

	if !token.Valid {
		log.Printf("token invalid")
		return false
	}

	return true
}
