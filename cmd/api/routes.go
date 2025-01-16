package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Encapsulate routing rules
// Keeps main function clean
// Easier to access test code by initializing an application instance and calling the routes() method on it.
func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance
	router := httprouter.New()

	// Register relevant methods, url patterns, and handler functions
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}
