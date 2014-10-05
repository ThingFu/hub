// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// DataSource for Things, States and Events
type DataSource interface {
	PutThing(*Thing) (Thing)

	GetThingCount() int
	GetThings() []Thing

	GetEventsCount() int
	GetEvents(int) []Event

	GetThingEvents(int, string) []Event

	CreateAppDB(*Thing) AppDB

	SaveThing(Thing)
	DeleteThing(Thing)

	PutEvent(*Event)
	SaveState(*Thing, map[string]interface{})
	ContainerAware

	SetEnvironment(Environment)
}
