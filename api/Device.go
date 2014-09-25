// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"fmt"
	bson "gopkg.in/mgo.v2/bson"
	"html/template"
	"io/ioutil"
	"log"
	"time"
)

// Representing a device or a thing
type Device struct {
	DatabaseId  bson.ObjectId `bson:"_id"`
	Id          string        `bson:"uid"`
	Name        string        `bson:"lbl"`
	Group       string        `bson:"grp"`
	Class       string        `bson:"c"`
	Type        string        `bson:"tid"`
	Enabled     bool          `bson:"en"`
	Description string        `bson:"desc"`
	LogEvents   bool          `bson:"logEvents"`
	Descriptor  DeviceType
	LastState   map[string]interface{} `bson:"state"`
	Attributes  []DeviceAttribute      `bson:"attrs"`
	LastEvent   time.Time
	LastCycle   time.Time
	Sensors     []Sensor `bson:"sub"`
	Content		string
}

func NewDevice() *Device {
	d := new(Device)

	return d
}

func (d Device) GetDatabaseId() string {
	return d.DatabaseId.String()
}

func (d *Device) UpdateLastCycle(t time.Time) {
	d.LastCycle = t
}

func (d *Device) GetId() string {
	return d.Id
}

func (d *Device) RenderWidget() template.HTML {
	path := d.Descriptor.Path

	fileContent, _ := ioutil.ReadFile(path + "/widget.html")
	stringContent := string(fileContent)

	tmpl, err := template.New("widget_" + d.Descriptor.Name).Parse(stringContent)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, d)

	htmlContent := buf.String()
	return template.HTML(htmlContent)
}

func (d Device) GetAttribute(name string) DeviceAttribute {
	for _, attr := range d.Attributes {
		if attr.Name == name {
			return attr
		}
	}
	return DeviceAttribute{}
}

func (d *Device) GetSensor(name string) *Sensor {
	for i, s := range d.Sensors {
		if s.Name == name {
			return &d.Sensors[i]
		}
	}
	return nil
}

func (d *Device) SaveAttribute(name string, value interface{}) error {
	found := false
	for i, attr := range d.Attributes {
		if attr.Name == name {
			d.Attributes[i].Value = value
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Attribute %s doesn't exist", name)
	}
	return nil
}

func (d *Device) UpdateLastEvent(t time.Time) {
	d.LastEvent = t
}
