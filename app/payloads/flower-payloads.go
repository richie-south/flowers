package payloads

import "net/http"

type WaterTimeline struct {
	Timestamp string `json:timestamp`
	Amount    string `json:amount`
}

// Flower type
type Flower struct {
	ID               int             `json:"id"`
	Title            string          `json:"title"`
	WateringTimeline []WaterTimeline `json:"waterTimeline"`
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
		fl.WateringTimeline = []WaterTimeline{
			{
				Timestamp: "hello",
				Amount:    "wello",
			},
		}
	}

	return nil
}
