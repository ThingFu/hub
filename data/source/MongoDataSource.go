// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

import (
	"fmt"
	"github.com/thingfu/hub/api"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"log"
)

type MongoDataSource struct {
	host      string
	env       api.Environment
	container api.Container
	session   *mgo.Session
}

func NewMongoDataSource() *MongoDataSource {
	m := new(MongoDataSource)

	session, err := mgo.Dial(m.host)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)
	m.session = session

	return m
}

func (m *MongoDataSource) GetThingCount() (count int) {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("devices").C("devices")
	count, _ = c.Count()

	fmt.Println(count)

	return
}

func (m *MongoDataSource) GetThings() []api.Thing {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("devices").C("devices")
	var results []api.Thing

	q := c.Find(bson.M{}).Sort("lbl")
	q.All(&results)

	return results
}

func (m *MongoDataSource) PutThing(t *api.Thing) api.Thing {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("devices").C("devices")
	t.DatabaseId = bson.NewObjectId()
	t.Descriptor = api.ThingType{}
	err := c.Insert(&t)
	if err != nil {
		log.Println(err)
	}

	return *t
}

func (m *MongoDataSource) GetThingEvents(limit int, id string) []api.Event {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("events").C("events")
	var results []api.Event
	if limit > 0 {
		c.Find(bson.M{"thing": id}).Limit(limit).Sort("-ts").All(&results)
	} else {
		c.Find(bson.M{"thing": id}).All(&results)
	}

	events := make([]api.Event, len(results))
	for i, v := range results {
		events[i] = v
	}
	return events
}

func (m *MongoDataSource) GetEventsCount() (count int) {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("events").C("events")
	count, _ = c.Count()

	return
}

func (m *MongoDataSource) GetEvents(limit int) []api.Event {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("events").C("events")
	var results []api.Event
	if limit > 0 {
		c.Find(bson.M{}).Limit(limit).Sort("-ts").All(&results)
	} else {
		c.Find(bson.M{}).All(&results)
	}

	events := make([]api.Event, len(results))
	for i, v := range results {
		events[i] = v
	}

	return events
}

func (m *MongoDataSource) SaveThing(dev api.Thing) {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("devices").C("devices")

	c.UpdateId(dev.GetDatabaseId(), &dev)
}

func (m *MongoDataSource) DeleteThing(dev api.Thing) {
	fmt.Println("Delete Thing")
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("devices").C("devices")
	fmt.Println(dev)
	err := c.RemoveId(dev.DatabaseId)
	if err != nil {
		fmt.Println("error deleting: ")
		fmt.Println(err)
	}
}

func (m *MongoDataSource) PutEvent(evt *api.Event) {
	go func() {
		session := m.session.Copy()
		defer session.Close()

		c := session.DB("events").C("events")

		c.Insert(evt)
	}()
}

func (m *MongoDataSource) SaveState(dev *api.Thing, state map[string]interface{}) {
	go func() {
		session := m.session.Copy()
		defer session.Close()

		c := session.DB("devices").C("devices")

		var change = mgo.Change{
			ReturnNew: true,
			Update: bson.M{
				"$set": bson.M{
					"state": state,
				},
			},
		}

		if _, err := c.FindId(dev.DatabaseId).Apply(change, &dev); err != nil {
			panic(err)
		}
	}()
}

func (s *MongoDataSource) CreateAppDB(t *api.Thing) api.AppDB {
	db := NewAppDB(t.Id)

	return db
}

func (s *MongoDataSource) ValidateWiring() {
	if s.env == nil {
		log.Fatal("Environment not wired to MongoDataSource")
	}
}

func (d *MongoDataSource) SetContainer(o api.Container) {
	d.container = o
}

func (d *MongoDataSource) SetEnvironment(o api.Environment) {
	d.env = o
}
