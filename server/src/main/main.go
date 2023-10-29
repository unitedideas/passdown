package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"passdown/server/src/main/db"
	"time"
)

func main() {
	// Create a new server and set timeout values.
	srv := &http.Server{
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	dsn := "username:password@tcp(127.0.0.1:3306)/your_db_name"
	database, err := db.NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	defer database.Connection.Close()

	router := CreateRouter(database)

	// Wrap the TreeMux router with your rate-limiting middleware using http.ServeMux
	mux := http.NewServeMux()
	mux.Handle("/", rateLimit(router)) // This should handle all your routes

	// Serving front-end static files
	fs := http.FileServer(http.Dir("./client/build"))
	mux.Handle("/static/", fs) // This should serve your React app's static files

	// Start the server in a goroutine so it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Channel to listen for interrupt signals to gracefully shutdown.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout before returning.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %+v", err)
	}

	log.Println("shutting down")
	os.Exit(0)
}
