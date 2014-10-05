// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Thing descriptors as defined in <home>/Things
type ThingType struct {
	TypeId            string `json:"tid"`
	LogEvents         bool   `json:"logEvents"`
	Name              string
	Description       string
	Protocol          string
	Vendor            string
	Model             string
	EventUpdateBuffer int
	Path              string
	Group             string
	Operations        []ThingOperation
	CycleTime         int
	Services     	  []ThingService
}

func (d *ThingType) GetService(name string) *ThingService {
	for i, s := range d.Services {
		if s.Name == name {
			return &d.Services[i]
		}
	}
	return nil
}
