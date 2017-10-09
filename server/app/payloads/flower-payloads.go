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

// Flower type
type Flower struct {
	ID                  bson.ObjectId       `json:"id" bson:"_id,omitempty"`
	Name                string              `json:"name"`
	FlowerType          string              `json:"flowerType" bson:"flowerType"`
	WaterIntervall      WaterIntervall      `json:"waterIntervall" bson:"waterIntervall"`
	WateringTimeline    []WaterTimelineItem `json:"waterTimeline" bson:"waterTimeline"`
	NextWateringSession time.Time           `json:"nextWateringSession" bson:"nextWateringSession"`
	Health              int                 `json:"helth"`
}

// RecivedWatering valid data server can recive when watering flower
type RecivedWatering struct {
	Amount string `json:"amount"`
}

// RecivedFlower valid data server can recive creating flower
type RecivedFlower struct {
	Name                     string    `json:"name"`
	FlowerType               string    `json:"flowerType"`
	OptimalWateringIntervall Intervall `json:"optimalWateringIntervall"`
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
	 * Calculate health of flower in a scale of 0 - 100 where 100 is best and 0 is worst
	 *  - start with water interval, if overdue lower health scale score,
	 *  - sircular graph
	 *
	 */
	// default values
	// check if NextWateringSession date has passed if has passed set date to now
	flower.Health = 100
	if len(flower.WateringTimeline) > 0 {
		//flower.Health = blaa(flower.WateringTimeline[len(flower.WateringTimeline)-1])
	}

	if flower.WateringTimeline == nil {
		flower.WateringTimeline = []WaterTimelineItem{}
	}

	return nil
}
