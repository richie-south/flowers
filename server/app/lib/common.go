package lib

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func WateringIntervallToText(intervall float64) string {
	days, hours := math.Modf(intervall)
	var dayString string
	var hourString string

	if int(days) > 1 {
		dayString = fmt.Sprintf("%v days", days)
	} else {
		dayString = fmt.Sprintf("%v day", days)
	}
	if int(hours*10) > 1 {
		hourString = fmt.Sprintf("%v hours", int(hours*10))
	} else {
		hourString = fmt.Sprintf("%v hour", int(hours*10))
	}

	return fmt.Sprintf("%v %v", dayString, hourString)
}

func DifferenceInDaysHoursToFloat(latest time.Time, second time.Time) (float64, error) {
	l := time.Date(latest.Year(), latest.Month(), latest.Day(), latest.Hour(), 0, 0, 0, time.UTC)
	s := time.Date(second.Year(), second.Month(), second.Day(), second.Hour(), 0, 0, 0, time.UTC)

	differenceInDays := int(l.Sub(s) / 24)
	differenceInHours := l.Sub(s)
	return strconv.ParseFloat(fmt.Sprintf("%d.%d", differenceInDays, differenceInHours), 64)
}
