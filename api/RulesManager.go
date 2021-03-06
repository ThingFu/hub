// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type RulesManager interface {
	GetRule(string) Rule
	GetRules() map[string]Rule
	Trigger(triggerType uint8, facts *RuleFacts)
	RegisterRule(Rule)

	ContainerAware
	SetThingManager(ThingManager)
	SetFactory(Factory)
}
