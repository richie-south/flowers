package payloads

import "fmt"

type Intervall struct {
	Days  int    `json:"days"`
	Hours int    `json:"hours"`
	Text  string `json:"text"`
}

type WaterIntervall struct {
	Optimal Intervall `json:"optimal"`
	/* 	OptimalText string    `json:"optimalText" bson:"optimalText"` */
	Current Intervall `json:"current"`
}

func (intervall *Intervall) ToText() string {
	var dayString string
	var hourString string

	dayString = fmt.Sprintf("%vd", intervall.Days)
	hourString = fmt.Sprintf("%vh", intervall.Hours)

	intervallText := fmt.Sprintf("%v %v", dayString, hourString)
	intervall.Text = intervallText
	return intervallText
}
