// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Attributes for a given Thing
type ThingAttribute struct {
	Name     string      `bson:"n"`
	Type     string      `bson:"t"`
	Config   bool        `bson:"cfg"`
	Required bool        `bson:"req"`
	Default  interface{} `bson:"def"`
	Value    interface{} `bson:"val"`
}

func (d *ThingAttribute) AsString() string {
	return d.Value.(string)
}

func (d *ThingAttribute) IsStringType() bool {
	if d.Type == "string" {
		return true
	}
	return false
}
