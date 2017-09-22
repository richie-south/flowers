package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/richie-south/flowers/app/handlers"
)

func NewFlowersRoute() http.Handler {
	router := chi.NewRouter()
	router.Route("/flowers", func(router chi.Router) {

		router.Get("/", handlers.ListFlowers) // GET /flowers
		router.Post("/", handlers.CreateFlower)

		// Subrouters
		router.Route("/{flowerID}", func(router chi.Router) {
			router.Use(handlers.FlowerToContext)
			router.Get("/", handlers.GetFlower) // GET /flowers/123

			// route: set latest watering time
			router.Route("/water", func(router chi.Router) {
				router.Post("/", handlers.WaterFlower)
			})
		})
	})
	return router
}
