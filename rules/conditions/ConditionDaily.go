// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conditions

import (
	"github.com/thingfu/hub/api"
	"time"
)

// Executed once a day
type Daily struct {
}

func (s Daily) Evaluate(when *api.RuleWhen, facts *api.RuleFacts, rule *api.Rule) bool {
	lastRun := rule.LastRun
	trigger, now := s.getTrigger()
	nowDiff := now.Sub(trigger).Seconds()

	if nowDiff > 0 {
		lastRunDiff := trigger.Sub(lastRun).Seconds()
		if lastRunDiff > 0 || lastRun.Year() == 0001 {
			return true
		}
	}
	return false
}

func (s Daily) getTrigger() (time.Time, time.Time) {
	now := time.Now()
	loc, _ := time.LoadLocation("Local")
	trigger := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	return trigger, now
}
