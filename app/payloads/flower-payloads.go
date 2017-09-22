package payloads

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type WaterTimeline struct {
	Timestamp string `json:"timestamp"`
	Amount    string `json:"amount"`
}

// Flower type
type Flower struct {
	ID               bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Name             string          `json:"name"`
	WateringTimeline []WaterTimeline `json:"waterTimeline"`
}

// RecivedFlower valid data server can recive
type RecivedFlower struct {
	Name string `json:"name"`
}

func (fl *Flower) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing
	/*
	 * metrics for:
	 * time since last watering
	 *
	 * merge watering timeline with sensor timeline
	 * gapfiller on timeline
	 * timeline of last watering or "event" sessions >--w---w--w--w->
	 *
	 *
	 */
	// default values
	if fl.WateringTimeline == nil {
		fl.WateringTimeline = []WaterTimeline{}
	}

	return nil
}
