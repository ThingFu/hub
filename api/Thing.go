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
	"regexp"
	"time"
)

// Representing a thing
type Thing struct {
	DatabaseId  bson.ObjectId                  `bson:"_id"`
	Id          string                         `bson:"uid"`
	Name        string                         `bson:"lbl"`
	Group       string                         `bson:"grp"`
	Class       string                         `bson:"c"`
	Type        string                         `bson:"tid"`
	Enabled     bool                           `bson:"en"`
	Description string                         `bson:"desc"`
	LogEvents   bool                           `bson:"logEvents"`
	Descriptor  ThingType                      `bson:",omitempty"`
	Attributes  map[string]ThingAttributeValue `bson:"attrs"`
	LastEvents  map[string]time.Time
	LastEvent   time.Time
	LastCycle   time.Time
	Data        map[string]interface{}  `bson:"data"`
	Content     string                  `bson:""`
	Services    map[string]ThingService `bson:"services"`
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

func (d Thing) GetAttributeValue(name string) ThingAttributeValue {
	if val, ok := d.Attributes[name]; ok {
		return val
	} else {
		return ThingAttributeValue{}
	}
}

func (d Thing) GetAttributeValues(expr string) (attrs map[string]ThingAttributeValue) {
	attrs = make(map[string]ThingAttributeValue)

	for k, v := range d.Attributes {
		match, _ := regexp.MatchString(expr, k)
		if match {
			attrs[k] = v
		}
	}

	return
}

//func (d *Thing) GetService(name string) *ThingService {
//	for i, s := range d.Services {
//		if s.Name == name {
//			return &d.Services[i]
//		}
//	}
//	return nil
//}

func (d *Thing) SaveAttributeValue(name string, value interface{}) error {
	if val, ok := d.Attributes[name]; ok {
		val.Value = value
		d.Attributes[name] = val
	} else {
		return fmt.Errorf("Attribute %s doesn't exist", name)
	}
	return nil
}

func (d *Thing) UpdateService(s *ThingService) {
	d.Services[s.Name] = *s
}

func (d *Thing) UpdateLastEvent(t time.Time) {
	d.LastEvent = t
}
