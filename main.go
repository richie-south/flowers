package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/richie-south/flowers/api/routes"
)

func main() {

	fmt.Println("Starting flowers")

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(middleware.Timeout(60 * time.Second))

	router.Mount("/api", routes.NewFlowersRoute())

	http.ListenAndServe(":3000", router)
}
