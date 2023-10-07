package main

import (
	"net/http"
	"passdown/server/src/db/model"
)

var shifts = []model.Shift{
	{1, 1, 1, "Machine 1 details", "Note 1", "2023-08-16T12:00:00Z"},
	// Add more sample shifts here
}

func main() {
	router := CreateRouter()

	// Wrap the TreeMux router with your rate-limiting middleware using http.ServeMux
	mux := http.NewServeMux()
	mux.Handle("/", rateLimit(router))

	http.ListenAndServe(":8080", mux)
}
