package cookies

import "net/http"

func SetCookie(w http.ResponseWriter, name string, value string) {
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
