package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole writes to console and moves to next middleware
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit the page")
		next.ServeHTTP(w, r)	
	})
}

// NoSurf Sets the base cookie to use when building a CSRF token cookie
func NoSurf(next http.Handler) http.Handler {
	csrfhandler := nosurf.New(next)

	csrfhandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfhandler
}