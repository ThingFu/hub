package tests

import (
	"github.com/thingfu/hub/api"
	"time"
)

func NewMockContainer() api.Container {
	c := new(MockContainer)
	c.thingManager = new(MockThingManager)

	return c
}

const (
	MIN  = 60
	HOUR = 60 * MIN
	DAY  = 24 * HOUR
	WEEK = 7 * DAY
)

func NewTime(when string) (t time.Time) {
	t = time.Now()
	switch {
	case when == "now":
		break

	case when == "yesterday":
		t = t.Add(-(time.Hour * 24))

	case when == "a_minute_Ago":
		t = t.Add(-(time.Minute))

	case when == "hour_ago":
		t = t.Add(-(time.Hour))

	case when == "2hours_ago":
		t = t.Add(-(time.Hour * 2))

	case when == "half_day":
		t = t.Add(-(time.Hour * 12))

	default:
		break
	}
	return
}
