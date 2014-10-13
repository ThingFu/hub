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
	Handle(*Thing, *ThingService, State)
	CreateThing(*Thing)
	RemoveThing(Thing)
	GetProtocolHandlerForThing(*Thing) (ProtocolHandler)

	LoadThings()

	Cycle()
	Actuate(*Thing, string, map[string]interface{})

	ContainerAware
	SetRulesManager(RulesManager)
	SetFactory(Factory)
	SetDataSource(DataSource)
	SetCommChannelManager(CommChannelManager)
}
