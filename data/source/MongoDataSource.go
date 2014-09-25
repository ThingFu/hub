// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

import (
	"github.com/go-home/hub/api"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"log"
	_ "fmt"
)

type MongoDataSource struct {
	host      		string
	env       		api.Environment
	container 		api.Container
	session			*mgo.Session
}

func NewMongoDataSource() *MongoDataSource {
	m := new (MongoDataSource)

	session, err := mgo.Dial(m.host)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)
	m.session = session

	return m
}

func (m *MongoDataSource) GetDevices() []api.Device {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("devices").C("devices")
	var results []api.Device

	q := c.Find(bson.M{}).Sort("lbl")
	q.All(&results)

	return results
}

func (m *MongoDataSource) PutDevice(dev *api.Device) {
	go func() {
		// TODO Asynchronous Operation
	}()
}

func (m *MongoDataSource) GetDeviceEventsCount() (count int) {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("events").C("events")
	count, _ = c.Count()

	return
}

func (m *MongoDataSource) GetDeviceEvents(limit int) []api.Event {
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

func (m *MongoDataSource) SaveDevice(dev api.Device) {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("devices").C("devices")

	c.UpdateId(dev.GetDatabaseId(), &dev)
}

func (m *MongoDataSource) PutEvent(evt *api.Event) {
	go func() {
		session := m.session.Copy()
		defer session.Close()

		c := session.DB("events").C("events")

		c.Insert(evt)
	}()
}

func (m *MongoDataSource )SaveState(dev *api.Device, state map[string] interface {}) {
	go func() {
		session := m.session.Copy()
		defer session.Close()

		c := session.DB("devices").C("devices")

		var change = mgo.Change {
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
/*
 var change = mgo.Change{
        ReturnNew: true,
        Update: bson.M{
            "$set": bson.M{
                "cp": time.Now(),
            }}}
    if changeInfo, err = collection.FindId(todo.Id).Apply(change, &todo); err != nil {
        panic(err)
    }
 */

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
