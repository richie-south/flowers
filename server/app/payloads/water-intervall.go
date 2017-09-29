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
	CurrentText string `json:"currentText" bson:"currentText"`
}

func (intervall *Intervall) ToText() string {
	var dayString string
	var hourString string

	if intervall.Days > 1 {
		dayString = fmt.Sprintf("%v days", intervall.Days)
	} else {
		dayString = fmt.Sprintf("%v day", intervall.Days)
	}
	if intervall.Hours > 1 {
		hourString = fmt.Sprintf("%v hours", intervall.Hours)
	} else {
		hourString = fmt.Sprintf("%v hour", intervall.Hours)
	}

	intervallText := fmt.Sprintf("%v %v", dayString, hourString)
	intervall.Text = intervallText
	return intervallText
}
