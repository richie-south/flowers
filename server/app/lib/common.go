package lib

import (
	"time"

	"github.com/richie-south/flowers/server/app/payloads"
)

func DifferenceInDaysHoursToIntervall(latest time.Time, second time.Time) payloads.Intervall {
	l := time.Date(latest.Year(), latest.Month(), latest.Day(), latest.Hour(), 0, 0, 0, time.UTC)
	s := time.Date(second.Year(), second.Month(), second.Day(), second.Hour(), 0, 0, 0, time.UTC)

	differenceInDays := int(l.Sub(s).Hours() / 24)
	differenceInHours := int(l.Sub(s).Hours())

	return payloads.Intervall{
		Days:  differenceInDays,
		Hours: differenceInHours,
	}
}
