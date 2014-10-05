// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"time"
)

const (
	EXECUTE       = 0
	NO_EXECUTE    = 1
	DEFER_EXECUTE = 2
)

const (
	TRIGGER_INTERVAL uint8 = 0
	TRIGGER_THING          = 1
)

type Rule struct {
	Id          string
	Path        string
	Name        string
	Description string
	Enabled     bool
	Priority    byte
	Async       bool
	Buffer      int
	LastRun     time.Time
	Targets     []string
	When        []RuleWhen
	Then        []RuleThen
}

type RuleFacts struct {
	Target string
	Thing  *Thing
	Service *ThingService
}

type RuleThen struct {
	Do     string
	Config map[string]interface{}
}

type RuleWhen struct {
	Target	string
	Service	string
	Event	string
}

type Consequence interface {
	Execute(config map[string]interface{}, container Container)
}

type Condition interface {
	Evaluate(when *RuleWhen, facts *RuleFacts, rule *Rule) bool
}
