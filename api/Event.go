// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/go-home/hub/utils"
	"time"
)

const (
	EVENT_SENSE = 0
)

type Event struct {
	Uid       string
	Device    string
	Sensor    string
	Ts        time.Time
	ShortText string
	LongText  string
	Event     int
	Data      interface{}
}

func NewEvent() *Event {
	evt := new(Event)
	evt.Ts = time.Now()

	return evt

}

func (e *Event) Ago() string {
	return utils.NewGoTime(e.Ts).Ago()
}
