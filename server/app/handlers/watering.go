package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/richie-south/flowers/server/app/services"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"github.com/richie-south/flowers/server/app/payloads"
)

func setNextWateringSession(OptimalWateringIntervall float64, timestamp time.Time) time.Time {
	intpart, floatpart := math.Modf(OptimalWateringIntervall)

	return timestamp.AddDate(0, 0, int(intpart)).Add(
		time.Hour * time.Duration(floatpart*10))
}

func WaterFlower(w http.ResponseWriter, req *http.Request) {
	flowerFromContext := req.Context().Value("flower").(*payloads.Flower)
	flowerID := fmt.Sprintf("%x", string(flowerFromContext.ID))

	decoder := json.NewDecoder(req.Body)
	waterAmout := payloads.RecivedWatering{}
	err := decoder.Decode(&waterAmout)
	if err != nil {
		render.Render(w, req, payloads.ErrUnexpected)
	}

	defer req.Body.Close()

	timestamp := time.Now()
	waterTimelineItem := &payloads.WaterTimelineItem{
		Timestamp: timestamp,
		Amount:    waterAmout.Amount,
	}

	err = dbAddToWaterTimeline(flowerID, *waterTimelineItem)
	if err != nil {
		render.Render(w, req, payloads.ErrWithDatabase)
		return
	}

	nextWateringSession := setNextWateringSession(
		flowerFromContext.OptimalWateringIntervall,
		timestamp,
	)

	flower, err := dbUpdateNextWateringSessionFlower(flowerID, nextWateringSession)
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
