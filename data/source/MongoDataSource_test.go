package source

import (
	"testing"
)

const (
	EVENTS_COUNT = 92
	THINGS_COUNT = 10

	THING_EVENTS_COUNT = 52
)

func TestThings(t *testing.T) {
	ds := NewMongoDataSource()

	count := ds.GetThingCount()
	if count != THINGS_COUNT {
		t.Errorf("Expected %d, returned %d", THINGS_COUNT, count)
	}

	things := ds.GetThings()
	count = len(things)
	if count != THINGS_COUNT {
		t.Errorf("Expected %d, returned %d", THINGS_COUNT, count)
	}

	count = len(ds.GetThingEvents(0, "nb-wt450-8"))
	if count != THING_EVENTS_COUNT {
		t.Errorf("Expected %d, returned %d", THING_EVENTS_COUNT, count)
	}

	count = len(ds.GetThingEvents(10, "nb-wt450-8"))
	if count != 10 {
		t.Errorf("Expected %d, returned %d", 10, count)
	}

	/*
		func (m *MongoDataSource) PutThing(dev *api.Thing) {
		func (m *MongoDataSource) SaveState(dev *api.Thing, state map[string]interface{}) {
	*/
}

func TestEvents(t *testing.T) {
	ds := NewMongoDataSource()

	events_count := ds.GetEventsCount()
	if events_count != EVENTS_COUNT {
		t.Errorf("Expected %d, returned %d", EVENTS_COUNT, events_count)
	}

	events_count = len(ds.GetEvents(0))
	if events_count != EVENTS_COUNT {
		t.Errorf("Expected %d, returned %d", EVENTS_COUNT, events_count)
	}

	events_count = len(ds.GetEvents(100))
	if events_count != EVENTS_COUNT {
		t.Errorf("Expected %d, returned %d", EVENTS_COUNT, events_count)
	}

	events_count = len(ds.GetEvents(50))
	if events_count != 50 {
		t.Errorf("Expected %d, returned %d", 50, events_count)
	}

	/*
		func (m *MongoDataSource) PutEvent(evt *api.Event) {
	*/
}
