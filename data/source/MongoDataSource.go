// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

import (
	"github.com/go-home/hub/api"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"log"
)

type MongoDataSource struct {
	host      string
	env       api.Environment
	container api.Container
}

func (m *MongoDataSource) GetDevices() []api.Device {
	session, err := mgo.Dial(m.host)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("devices").C("devices")
	var results []api.Device

	err = c.Find(bson.M{}).All(&results)

	return results
}

func (m *MongoDataSource) PutDevice(dev *api.Device) {
	go func() {
		// TODO Asynchronous Operation
	}()
}

func (m *MongoDataSource) GetDeviceEvents(limit int) []api.Event {
	session, err := mgo.Dial(m.host)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("events").C("events")
	var results []api.Event
	if limit > 0 {
		err = c.Find(bson.M{}).Limit(limit).Sort("-ts").All(&results)
	} else {
		err = c.Find(bson.M{}).All(&results)
	}

	events := make([]api.Event, len(results))
	for i, v := range results {
		events[i] = v
	}

	return events
}

func (m *MongoDataSource) SaveDevice(dev api.Device) {
	session, err := mgo.Dial(m.host)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("devices").C("devices")

	c.UpdateId(dev.GetDatabaseId(), &dev)
}

func (m *MongoDataSource) PutEvent(evt *api.Event) {
	go func() {
		session, err := mgo.Dial(m.host)
		defer session.Close()
		if err != nil {
			log.Fatal(err)
		}
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("events").C("events")

		c.Insert(evt)
	}()
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
