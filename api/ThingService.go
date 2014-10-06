package api

import "time"

type ThingService struct {
	Name      string `json:"name"`
	Label     string `json:"lbl"`
	LastEvent time.Time
}

func (d *ThingService) UpdateLastEvent(t time.Time) {
	d.LastEvent = t
}
