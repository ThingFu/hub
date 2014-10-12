// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/thingfu/hub/api"
)

type AdapterWeather struct {
}

func (d *AdapterWeather) OnWrite(t *api.Thing, op string, params api.WriteRequest, db api.AppDB) {

}

func (d *AdapterWeather) Cycle(dev *api.Thing) {

}

func (d *AdapterWeather) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest) (state map[string]interface{}) {
	return nil
}

func (d *AdapterWeather) GetEventText(dev *api.Thing, sensor *api.ThingService, state map[string]interface{}) (shortText string, longText string) {
	return
}
