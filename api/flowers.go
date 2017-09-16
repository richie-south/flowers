package api

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// FlowerCtx parses id and finds flower object
func FlowerCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flowerIDString := chi.URLParam(r, "flowerID")
		flowerID, err := strconv.Atoi(flowerIDString)

		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		flower, err := dbGetFlower(flowerID)
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "flower", flower)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ListFlowers(w http.ResponseWriter, r *http.Request) {
	flowerList := []render.Renderer{}
	for _, flower := range flowers {
		flowerList = append(flowerList, flower)
	}

	if err := render.RenderList(w, r, flowerList); err != nil {
		render.Render(w, r, ErrUnexpected)
		return
	}
}

func GetFlower(w http.ResponseWriter, r *http.Request) {
	flower := r.Context().Value("flower").(*Flower)
	if err := render.Render(w, r, flower); err != nil {
		render.Render(w, r, ErrUnexpected)
		return
	}
}

var flowers = []*Flower{
	{ID: 1, Title: "Flower 1"},
	{ID: 2, Title: "Flower 2"},
	{ID: 3, Title: "Flower 3"},
	{ID: 4, Title: "Flower 4"},
	{ID: 5, Title: "Flower 5"},
}

// only testing implemnt real db later
func dbGetFlower(id int) (*Flower, error) {
	for _, f := range flowers {
		if f.ID == id {
			return f, nil
		}
	}
	return nil, errors.New("flower not found")
}
