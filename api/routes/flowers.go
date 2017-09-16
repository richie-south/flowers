package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/richie-south/flowers/api"
)

func NewFlowersRoute() http.Handler {
	router := chi.NewRouter()
	router.Route("/flowers", func(router chi.Router) {

		router.Get("/", api.ListFlowers) // GET /flowers

		// Subrouters
		router.Route("/{flowerID}", func(router chi.Router) {
			router.Use(api.FlowerCtx)
			router.Get("/", api.GetFlower) // GET /flowers/123
		})
	})
	return router
}
