// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type DeviceAttribute struct {
	Name     string	`bson:"n"`
	Type     string	`bson:"t"`
	Config   bool	`bson:"cfg"`
	Required bool	`bson:"req"`
	Default  interface{}	`bson:"def"`
	Value    interface{}	`bson:"val"`
}

func (d *DeviceAttribute) AsString() string {
	return d.Value.(string)
}

func (d *DeviceAttribute) IsStringType() bool {
	if d.Type == "string" {
		return true
	}
	return false
}
