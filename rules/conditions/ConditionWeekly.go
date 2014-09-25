// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conditions

import (
	"github.com/go-home/hub/api"
	"time"
)

// Executed once a week
type Weekly struct {
}

func (s Weekly) Evaluate(when *api.RuleWhen, facts *api.RuleFacts, rule *api.Rule) bool {
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

func (s Weekly) getTrigger() (time.Time, time.Time) {
	now := time.Now()

	// weekday := now.Weekday()

	loc, _ := time.LoadLocation("Local")
	trigger := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc)

	return trigger, now
}

/*
exports.name = "weekly";
var moment = require("moment");

exports.fn = function ($when, $fact, $rule) {
    var lastRun = $rule.lastRun;
    var now = moment();
    var triggerWhen = moment().startOf('week');

    if (now.diff(triggerWhen) > 0) {
        if (lastRun == undefined || triggerWhen.diff(lastRun) > 0) {
            return true;
        }
    }
    return false;
}
*/
