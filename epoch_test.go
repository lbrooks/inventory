package inventory_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/lbrooks/inventory"
)

const layoutISO = "2006-01-02"

var tests = []struct {
	date  time.Time
	epoch int
}{
	{time.Date(1969, time.December, 31, 0, 0, 0, 0, time.UTC), -1},
	{time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC), 0},
	{time.Date(1970, time.January, 2, 0, 0, 0, 0, time.UTC), 1},
	{time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC), 9},
	{time.Date(2022, time.May, 3, 0, 0, 0, 0, time.UTC), 19115},
}

func TestDateToEpoch(t *testing.T) {
	for _, v := range tests {
		t.Run(v.date.Format(layoutISO), func(t *testing.T) {
			epoch := inventory.DateToEpoch(v.date)
			if epoch != v.epoch {
				t.Errorf("DateToEpoch(%s) = %d; want %d", v.date.Format(layoutISO), epoch, v.epoch)
			}
		})
	}
}

func BenchmarkDateToEpoch(b *testing.B) {
	for _, v := range tests {
		b.Run(v.date.Format(layoutISO), func(b *testing.B) {
			epoch := inventory.DateToEpoch(v.date)
			if epoch != v.epoch {
				b.Errorf("DateToEpoch(%s) = %d; want %d", v.date.Format(layoutISO), epoch, v.epoch)
			}
		})
	}
}

func TestEpochToDate(t *testing.T) {
	for _, v := range tests {
		t.Run(strconv.Itoa(v.epoch), func(t *testing.T) {
			day := inventory.EpochToDate(v.epoch)
			if day != v.date {
				t.Errorf("EpochToDate(%s) = %s; want %s", strconv.Itoa(v.epoch), day, v.date)
			}
		})
	}
}

func BenchmarkEpochToDate(b *testing.B) {
	for _, v := range tests {
		b.Run(strconv.Itoa(v.epoch), func(b *testing.B) {
			day := inventory.EpochToDate(v.epoch)
			if day != v.date {
				b.Errorf("EpochToDate(%s) = %s; want %s", strconv.Itoa(v.epoch), day, v.date)
			}
		})
	}
}
