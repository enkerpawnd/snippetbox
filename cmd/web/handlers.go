package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Can be used add a custom header, where first arg
	// is header name, second arg is header value
	w.Header().Add("Server", "Go") // must be called before a writeHeader
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// sanitizing input and for our specific reason we want
	// to ensure that id is a positive value greater than 0
	// so we can "whitelist" what is valid
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a specific snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet"))
}
