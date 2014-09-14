// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type RulesService interface {
	GetRules() map[string]Rule
	Trigger(triggerType int, facts *RuleFacts)
	ContainerAware
	RegisterRule(Rule)

	// Injection
	SetDeviceService(DeviceService)
	SetFactory(Factory)
}
