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

	default:
		break
	}
	return
}
