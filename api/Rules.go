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
	TRIGGER_INTERVAL = 0
	TRIGGER_DEVICE   = 1
)

type Rule struct {
	Id          string
	Path        string
	Name        string
	Description string
	Enabled     bool
	Priority    int
	Targets     []string
	When        []RuleWhen
	Then        []RuleThen
	LastRun     time.Time
	Async       bool
	Buffer      int
}

type RuleFacts struct {
	Target string
	Device *Device
}

type RuleThen struct {
	Do     string
	Config map[string]interface{}
}

type RuleWhen struct {
	Target  string
	Trigger string
}

type Consequence interface {
	Execute(config map[string]interface{}, container Container)
}

type Condition interface {
	Evaluate(when *RuleWhen, facts *RuleFacts, rule *Rule) bool
}