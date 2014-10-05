// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Service for handling Things
type ThingManager interface {
	GetThing(string) (Thing, bool)
	SaveThing(Thing)
	GetThingType(string) (ThingType, error)
	GetThingTypes() map[string]ThingType
	RegisterThingType(ThingType)
	RegisterThing(Thing)
	GetThings() []Thing
	Handle(*Thing, *ThingService, map[string]interface{})
	CreateThing(*Thing)
	RemoveThing(Thing)

	LoadThings()

	Cycle()
	Actuate(t *Thing, op string, params map[string]interface{})

	ContainerAware
	SetRulesManager(RulesManager)
	SetFactory(Factory)
	SetDataSource(DataSource)
}
