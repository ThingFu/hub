// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Attributes for a given Thing
type ThingAttribute struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Config   bool        `json:"config"`
	Required bool        `json:"required"`
	Default  interface{} `json:"default"`
}

func (d *ThingAttribute) IsStringType() bool {
	if d.Type == "string" {
		return true
	}
	return false
}
