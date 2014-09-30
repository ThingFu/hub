package api

import "time"

type ThingService struct {
	Name      string `bson:"n"`
	Label     string `bson:"lbl"`
	Code      string `bson:"code"`
	LastEvent time.Time
}

func (d *ThingService) UpdateLastEvent(t time.Time) {
	d.LastEvent = t
}
