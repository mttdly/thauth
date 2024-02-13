package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// make a new router instance
	r := chi.NewRouter()

	// register a middleware for logging
	r.Use(middleware.Logger)

	// handler for requests
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("root"))
	})

	// starts up the server instance on port 3000
	http.ListenAndServe(":3000", r)
}
