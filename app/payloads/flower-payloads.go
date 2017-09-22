package payloads

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type WaterTimelineItem struct {
	Timestamp time.Time `json:"timestamp"`
	Amount    string    `json:"amount"`
}

type RecivedWatering struct {
	Amount string `json:"amount"`
}

// Flower type
type Flower struct {
	ID                       bson.ObjectId       `json:"id" bson:"_id,omitempty"`
	Name                     string              `json:"name"`
	FlowerType               string              `json:"flowerType"`
	OptimalWateringIntervall int                 `json:"optimalWateringIntervall"`
	WateringTimeline         []WaterTimelineItem `json:"waterTimeline" bson:"waterTimeline"`
}

// RecivedFlower valid data server can recive
type RecivedFlower struct {
	Name                     string `json:"name"`
	FlowerType               string `json:"flowerType"`
	OptimalWateringIntervall int    `json:"optimalWateringIntervall"`
}

func (flower *Flower) Render(w http.ResponseWriter, r *http.Request) error {
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
	if flower.WateringTimeline == nil {
		flower.WateringTimeline = []WaterTimelineItem{}
	}

	return nil
}
