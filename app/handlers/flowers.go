package handlers

import (
	"context"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/richie-south/flowers/app/payloads"
	"github.com/richie-south/flowers/app/services"
)

// FlowerCtx parses id and finds flower object
func FlowerToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flowerID := chi.URLParam(r, "flowerID")

		flower, err := dbGetFlower(flowerID)
		if err != nil {
			render.Render(w, r, payloads.ErrNotFound)
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

	err := render.RenderList(w, r, flowerList)
	if err != nil {
		render.Render(w, r, payloads.ErrUnexpected)
		return
	}
}

func GetFlower(w http.ResponseWriter, r *http.Request) {
	flower := r.Context().Value("flower").(*payloads.Flower)
	err := render.Render(w, r, flower)
	if err != nil {
		render.Render(w, r, payloads.ErrUnexpected)
		return
	}
}

var flowers = []*payloads.Flower{
	{ID: 1, Title: "Flower 1"},
	{ID: 2, Title: "Flower 2"},
	{ID: 3, Title: "Flower 3"},
	{ID: 4, Title: "Flower 4"},
	{ID: 5, Title: "Flower 5"},
	{
		ID:    6,
		Title: "Flower 6",
		WateringTimeline: []payloads.WaterTimeline{
			{
				Timestamp: "asd",
				Amount:    "small",
			},
		},
	},
}

// only testing implemnt real db later
func dbGetFlower(id string) (payloads.Flower, error) {
	result := payloads.Flower{}

	query := func(collection *mgo.Collection) error {
		return collection.FindId(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	}

	err := services.WithCollection("flower", query)
	if err != nil {
		panic("Error with collection pool")
	}
	return result, err
}
