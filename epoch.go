package inventory

import (
	"sync"
	"time"
)

var once sync.Once
var epoch time.Time

func getEpoch() time.Time {
	once.Do(func() {
		epoch = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	})

	return epoch
}

func DateToEpoch(date time.Time) int {
	current := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	epoch := getEpoch()

	return int(current.Sub(epoch).Hours() / 24)
}

func EpochToDate(epochDays int) time.Time {
	return getEpoch().AddDate(0, 0, epochDays)
}
