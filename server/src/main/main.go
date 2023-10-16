package main

import (
	"net/http"
)

func main() {
	router := CreateRouter()

	// Wrap the TreeMux router with your rate-limiting middleware using http.ServeMux
	mux := http.NewServeMux()
	mux.Handle("/", rateLimit(router)) // This should handle all your routes

	// Serving front-end static files
	fs := http.FileServer(http.Dir("./client/build"))
	mux.Handle("/static/", fs) // This should serve your React app's static files

	http.ListenAndServe(":8080", mux)
}
