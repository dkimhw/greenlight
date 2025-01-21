package main

import (
	"net/http"
)

// Sends back info about app status, os, and version.
// This is a method on the application struct. This is an effective way to
// expose dependencies available to the handlers without resorting to global variables.
// Any dependency that this method requires can simply be includd as a field in the application
// struct.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map which holds the information that we want to send in the response.
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}
