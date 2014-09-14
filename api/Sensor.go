package api

import "time"

type Sensor struct {
	Name      string `bson:"n"`
	Label     string `bson:"lbl"`
	Code      string `bson:"code"`
	LastEvent time.Time
}

func (d *Sensor) UpdateLastEvent(t time.Time) {
	d.LastEvent = t
}
