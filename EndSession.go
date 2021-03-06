package main

import (
	"net/http"
)

/*
	Handles the request to EndSession.
	Ends the session associated with the client
	if there is one.
	Returns an empty response.
*/

func EndSession(w http.ResponseWriter, r *http.Request) {
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete the session
	session.Options.MaxAge = -1

	if session.Save(r, w) != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
