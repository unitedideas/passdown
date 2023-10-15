package main

import (
	"net/http"
	"passdown/server/src/main/handler"

	"github.com/dimfeld/httptreemux/v5"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5) // 1 request per second, with a burst of 5

func rateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Exclude your IP from rate limiting
		if r.RemoteAddr == "2601:601:510:2971:54ef:589c:eddc:8154" {
			next.ServeHTTP(w, r)
			return
		}

		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CreateRouter() *httptreemux.TreeMux {
	router := httptreemux.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		w.Write([]byte("Hello, world! - Did this change?"))
	})

	// Shift related routes
	router.GET("/api/v1/shifts", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		handler.HandleGetShiftsRequest(w, r)
	})
	// Add more routes related to 'shift' here

	// Redirect all other GET requests back to react-router (or a custom 404 handler)
	router.NotFoundHandler = http.HandlerFunc(handlePageNotFound)

	return router
}

func handlePageNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
