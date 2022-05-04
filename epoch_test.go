package inventory_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/lbrooks/inventory"
)

var tests = []struct {
	isoDate string
	date    time.Time
	epoch   int
}{
	{"1969-12-31", time.Date(1969, time.December, 31, 0, 0, 0, 0, time.UTC), -1},
	{"1970-01-01", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC), 0},
	{"1970-01-02", time.Date(1970, time.January, 2, 0, 0, 0, 0, time.UTC), 1},
	{"1970-01-10", time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC), 9},
	{"2022-05-03", time.Date(2022, time.May, 3, 0, 0, 0, 0, time.UTC), 19115},
}

func TestDateToEpoch(t *testing.T) {
	for _, v := range tests {
		t.Run(v.isoDate, func(t *testing.T) {
			epoch := inventory.DateToEpoch(v.date)
			if epoch != v.epoch {
				t.Errorf("DateToEpoch(%s) = %d; want %d", v.isoDate, epoch, v.epoch)
			}
		})
	}
}

func BenchmarkDateToEpoch(b *testing.B) {
	for _, v := range tests {
		b.Run(v.isoDate, func(b *testing.B) {
			epoch := inventory.DateToEpoch(v.date)
			if epoch != v.epoch {
				b.Errorf("DateToEpoch(%s) = %d; want %d", v.isoDate, epoch, v.epoch)
			}
		})
	}
}

func TestIsoDateToEpoch(t *testing.T) {
	for _, v := range tests {
		t.Run(v.isoDate, func(t *testing.T) {
			epoch, err := inventory.IsoDateToEpoch(v.isoDate)
			if err != nil {
				t.Error(err)
			}
			if epoch != v.epoch {
				t.Errorf("IsoDateToEpoch(%s) = %d; want %d", v.isoDate, epoch, v.epoch)
			}
		})
	}
}

func BenchmarkIsoDateToEpoch(b *testing.B) {
	for _, v := range tests {
		b.Run(v.isoDate, func(b *testing.B) {
			epoch, err := inventory.IsoDateToEpoch(v.isoDate)
			if err != nil {
				b.Error(err)
			}
			if epoch != v.epoch {
				b.Errorf("IsoDateToEpoch(%s) = %d; want %d", v.isoDate, epoch, v.epoch)
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

func TestEpochToIsoDate(t *testing.T) {
	for _, v := range tests {
		t.Run(strconv.Itoa(v.epoch), func(t *testing.T) {
			day := inventory.EpochToIsoDate(v.epoch)
			if day != v.isoDate {
				t.Errorf("EpochToIsoDate(%s) = %s; want %s", strconv.Itoa(v.epoch), day, v.isoDate)
			}
		})
	}
}

func BenchmarkEpochToIsoDate(b *testing.B) {
	for _, v := range tests {
		b.Run(strconv.Itoa(v.epoch), func(b *testing.B) {
			day := inventory.EpochToIsoDate(v.epoch)
			if day != v.isoDate {
				b.Errorf("EpochToIsoDate(%s) = %s; want %s", strconv.Itoa(v.epoch), day, v.isoDate)
			}
		})
	}
}
