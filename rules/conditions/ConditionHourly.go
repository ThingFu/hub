// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conditions

import (
	"fmt"
	"github.com/thingfu/hub/api"
	"time"
)

// Executed every hour
type Hourly struct {
}

func (s Hourly) Evaluate(when *api.RuleWhen, facts *api.RuleFacts, rule *api.Rule) bool {
	fmt.Println("----------- ConditionHourly ---------------------")
	lastRun := rule.LastRun
	trigger, now := s.getTrigger()
	nowDiff := now.Sub(trigger).Seconds()
	fmt.Println("LastRun:")
	fmt.Println(lastRun)
	fmt.Println("Now: ")
	fmt.Println(now)
	fmt.Println("Diff: ")
	fmt.Println(nowDiff)
	fmt.Println("----------- ConditionHourly ---------------------")

	if nowDiff > 0 {
		lastRunDiff := trigger.Sub(lastRun).Seconds()
		if lastRunDiff > 0 || lastRun.Year() == 0001 {
			return true
		}
	}
	return false
}

func (s Hourly) getTrigger() (time.Time, time.Time) {
	now := time.Now()
	loc, _ := time.LoadLocation("Local")
	trigger := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, loc)

	return trigger, now
}
