package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("Starting flowers")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	// REST routes for "flowers"
	r.Route("/flowers", func(r chi.Router) {

		r.Post("/", createFlower) // POST /flowers
		r.Get("/", listFlowers)   // GET /flowers

		// Subrouters
		r.Route("/{flowerID}", func(r chi.Router) {
			// r.Use(FlowerCtx)
			r.Get("/", getFlower) // GET /flowers/123
		})
	})

	http.ListenAndServe(":3333", r)
}

func createFlower(w http.ResponseWriter, r *http.Request) {
}

func listFlowers(w http.ResponseWriter, r *http.Request) {
}

func getFlower(w http.ResponseWriter, r *http.Request) {
}
