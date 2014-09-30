// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Service for handling Things
type ThingService interface {
	GetThing(string) (Thing, bool)
	SaveThing(Thing)
	GetThingType(string) ThingType
	GetThingTypes() map[string]ThingType
	RegisterThingType(ThingType)
	RegisterThing(Thing)
	GetThings() []Thing
	Handle(*Thing, *Sensor, map[string]interface{})

	Cycle()
	Actuate(t *Thing, op string, params map[string]interface{})

	ContainerAware
	SetRulesService(RulesService)
	SetFactory(Factory)
	SetDataSource(DataSource)
}
