// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type DeviceType struct {
	TypeId            string `json:"tid"`
	Name              string
	Description       string
	Protocol          string
	Vendor            string
	Model             string
	EventUpdateBuffer int
	Path              string
	Group             string
	Operations        []DeviceOperation
	CycleTime         int
}
