package inventory

import (
	"sync"
	"time"
)

const layoutISO = "2006-01-02"

var once sync.Once
var epoch time.Time

func getEpoch() time.Time {
	once.Do(func() {
		epoch = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	})

	return epoch
}

func daysBetween(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

func DateToEpoch(date time.Time) int {
	return daysBetween(getEpoch(),
		time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC))
}

func IsoDateToEpoch(date string) (int, error) {
	current, err := time.Parse(layoutISO, date)
	if err != nil {
		return -1, err
	}
	return daysBetween(getEpoch(), current), nil
}

func EpochToDate(epochDays int) time.Time {
	return getEpoch().AddDate(0, 0, epochDays)
}

func EpochToIsoDate(epochDays int) string {
	return EpochToDate(epochDays).Format(layoutISO)
}
