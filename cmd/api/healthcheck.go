package main

import (
	"fmt"
	"net/http"
)

// Sends back info about app status, os, and version.
// This is a method on the application struct. This is an effective way to
// expose dependencies available to the handlers without resorting to global variables.
// Any dependency that this method requires can simply be includd as a field in the application
// struct.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
