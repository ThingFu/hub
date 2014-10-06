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

// Representing a thing
type Thing struct {
	DatabaseId  bson.ObjectId             `bson:"_id"`
	Id          string                    `bson:"uid"`
	Name        string                    `bson:"lbl"`
	Group       string                    `bson:"grp"`
	Class       string                    `bson:"c"`
	Type        string                    `bson:"tid"`
	Enabled     bool                      `bson:"en"`
	Description string                    `bson:"desc"`
	LogEvents   bool                      `bson:"logEvents"`
	Descriptor  ThingType                 `bson:",omitempty"`
	LastState   map[string]interface{}    `bson:"state"`
	Attributes  map[string]ThingAttribute `bson:"attrs"`
	LastEvent   time.Time
	LastCycle   time.Time
	Data        map[string]interface{} `bson:"data"`
	Content     string
}

func NewThing() *Thing {
	d := new(Thing)

	return d
}

func (d Thing) GetDatabaseId() string {
	return d.DatabaseId.String()
}

func (d *Thing) UpdateLastCycle(t time.Time) {
	d.LastCycle = t
}

func (d *Thing) GetId() string {
	return d.Id
}

func (d *Thing) RenderWidget() template.HTML {
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

func (d Thing) GetAttribute(name string) ThingAttribute {
	if val, ok := d.Attributes[name]; ok {
		return val
	} else {
		return ThingAttribute{}
	}
}

//func (d *Thing) GetService(name string) *ThingService {
//	for i, s := range d.Services {
//		if s.Name == name {
//			return &d.Services[i]
//		}
//	}
//	return nil
//}

func (d *Thing) SaveAttribute(name string, value interface{}) error {
	if val, ok := d.Attributes[name]; ok {
		val.Value = value
		d.Attributes[name] = val
	} else {
		return fmt.Errorf("Attribute %s doesn't exist", name)
	}
	return nil
}

func (d *Thing) UpdateLastEvent(t time.Time) {
	d.LastEvent = t
}

// Things
type Things []Thing

func (slice Things) Len() int {
	return len(slice)
}

func (slice Things) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice Things) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
