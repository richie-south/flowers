package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/richie-south/flowers/server/app/services"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"github.com/richie-south/flowers/server/app/payloads"
)

func WaterFlower(w http.ResponseWriter, req *http.Request) {
	flowerID := chi.URLParam(req, "flowerID")
	decoder := json.NewDecoder(req.Body)
	waterAmout := payloads.RecivedWatering{}
	err := decoder.Decode(&waterAmout)
	if err != nil {
		render.Render(w, req, payloads.ErrUnexpected)
	}

	defer req.Body.Close()

	waterTimelineItem := &payloads.WaterTimelineItem{
		Timestamp: time.Now(),
		Amount:    waterAmout.Amount,
	}

	err = dbAddToWaterTimeline(flowerID, *waterTimelineItem)
	if err != nil {
		render.Render(w, req, payloads.ErrWithDatabase)
		return
	}

	flower, err := dbGetFlower(flowerID)
	if err != nil {
		render.Render(w, req, payloads.ErrWithDatabase)
		return
	}

	render.Render(w, req, &flower)
}

func dbAddToWaterTimeline(id string, waterTime payloads.WaterTimelineItem) error {
	query := func(collection *mgo.Collection) error {
		if bson.IsObjectIdHex(id) {
			err := collection.Update(
				bson.M{"_id": bson.ObjectIdHex(id)},
				bson.M{"$push": bson.M{"waterTimeline": waterTime}},
			)

			if err != nil {
				fmt.Println("error", err)
				return err
			}

			return nil
		}

		return errors.New("Id not object hex")
	}

	err := services.WithCollection("flower", query)
	if err != nil {
		return err
	}

	return nil
}
