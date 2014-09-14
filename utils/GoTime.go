// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"time"
)

const (
	ONE_MINUTE  = 60
	TWO_MINUTES = 120
	ONE_HOUR    = ONE_MINUTE * 60
	TWO_HOURS   = ONE_HOUR * 2
	ONE_DAY     = ONE_HOUR * 24
	TWO_DAYS    = ONE_DAY * 2
	ONE_MONTH   = ONE_DAY * 30
	TWO_MONTHS  = ONE_MONTH * 2
	ONE_YEAR    = ONE_DAY * 365
)

type GoTime struct {
	_t time.Time
}

func (t *GoTime) SetTime(time time.Time) {
	t._t = time
}

func (t GoTime) Ago() string {
	since := int(time.Since(t._t).Seconds())

	if since < ONE_MINUTE {
		return fmt.Sprintf("%d secs", since)
	} else if since >= ONE_MINUTE && since < TWO_MINUTES {
		return "a minute"
	} else if since > TWO_MINUTES {
		if since < ONE_HOUR {
			return fmt.Sprintf("%d mins", since/60)
		} else if since >= ONE_HOUR && since < TWO_HOURS {
			return fmt.Sprintf("1 hour")
		} else if since <= ONE_DAY {
			return fmt.Sprintf("%d hours", since/60/60)
		} else if since > ONE_DAY && since < TWO_DAYS {
			return fmt.Sprintf("1 day")
		} else if since <= ONE_MONTH {
			return fmt.Sprintf("%d days", since/60/60/24)
		} else if since > ONE_MONTH && since < TWO_MONTHS {
			return fmt.Sprintf("a month")
		} else {
			return fmt.Sprintf("%d months", since/60/60/24/30)
		}
	}
	return fmt.Sprintf("%d", since)
}

func (t GoTime) GetTime() time.Time {
	return t._t
}

func (t GoTime) Elapsed(secs int) bool {
	then := t.GetTime()
	now := time.Now()

	d := int(now.Sub(then).Seconds())

	if d > secs {
		return true
	} else {
		return false
	}
}

/*
	Few seconds ago
	if seconds < 60
		x seconds ago
	if seconds > 60 < 120
		a minute ago
	if seconds >= 120
		x minutes ago
	if minutes > 60 < 120
		an hour ago
	if minutes > 120 < 2880
		x hours ago
	if minutes > 2880


*/
