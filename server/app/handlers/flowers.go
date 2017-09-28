package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/richie-south/flowers/server/app/lib"
	"github.com/richie-south/flowers/server/app/payloads"
	"github.com/richie-south/flowers/server/app/services"
)

// FlowerToContext parses id and finds flower object
func FlowerToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flowerID := chi.URLParam(r, "flowerID")

		flower, err := dbGetFlower(flowerID)
		if err != nil {
			render.Render(w, r, payloads.ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "flower", &flower)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ListFlowers(w http.ResponseWriter, r *http.Request) {

	flowers, err := dbGetFlowers()
	if err != nil {
		render.Render(w, r, payloads.ErrWithDatabase)
		return
	}

	flowerList := []render.Renderer{}
	for _, flower := range flowers {
		flowerList = append(flowerList, flower)
	}

	err = render.RenderList(w, r, flowerList)
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

func CreateFlower(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	f := payloads.RecivedFlower{}
	err := decoder.Decode(&f)
	if err != nil {
		render.Render(w, req, payloads.ErrUnexpected)
		return
	}
	defer req.Body.Close()

	id := bson.NewObjectId()
	flower := &payloads.Flower{
		ID:         id,
		Name:       f.Name,
		FlowerType: f.FlowerType,
		WaterIntervall: payloads.WaterIntervall{
			Optimal:     f.OptimalWateringIntervall,
			OptimalText: lib.WateringIntervallToText(f.OptimalWateringIntervall),
			CurrentText: "",
		},
		NextWateringSession: time.Now(),
	}

	err = dbInsertFlower(*flower)

	if err != nil {
		render.Render(w, req, payloads.ErrWithDatabase)
		return
	}

	render.Render(w, req, flower)
}

func dbGetFlowers() ([]*payloads.Flower, error) {
	result := []*payloads.Flower{}

	query := func(collection *mgo.Collection) error {
		return collection.Find(nil).All(&result)
	}

	err := services.WithCollection("flower", query)
	if err != nil {
		return result, err
	}

	return result, nil
}

func dbGetFlower(id string) (payloads.Flower, error) {
	result := payloads.Flower{}

	query := func(collection *mgo.Collection) error {
		if bson.IsObjectIdHex(id) {
			return collection.FindId(bson.ObjectIdHex(id)).One(&result)
		}
		return errors.New("Id not object hex")
	}

	err := services.WithCollection("flower", query)
	if err != nil {
		return result, err
	}

	return result, err
}

func dbInsertFlower(flower payloads.Flower) error {
	query := func(collection *mgo.Collection) error {
		err := collection.Insert(flower)
		if err != nil {
			return err
		}

		return nil
	}

	err := services.WithCollection("flower", query)
	if err != nil {
		return err
	}

	return nil
}

func dbUpdateNextWateringSessionFlower(id string, nextWateringSession time.Time) error {
	query := func(collection *mgo.Collection) error {
		if bson.IsObjectIdHex(id) {
			err := collection.Update(
				bson.M{"_id": bson.ObjectIdHex(id)},
				bson.M{"$set": bson.M{"nextWateringSession": nextWateringSession}},
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

func dbUpdateCurrentWaterIntervall(id string, intervall string) error {
	query := func(collection *mgo.Collection) error {
		if bson.IsObjectIdHex(id) {
			err := collection.Update(
				bson.M{"_id": bson.ObjectIdHex(id)},
				bson.M{"$set": bson.M{"waterIntervall.currentText": intervall}},
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
