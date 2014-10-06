// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/thingfu/hub/api"
)

type AdapterKnockSensor433 struct {
}

func (o *AdapterKnockSensor433) Cycle(dev *api.Thing) {

}

func (d *AdapterKnockSensor433) OnActuate(t *api.Thing, op string, params map[string]interface{}, db api.AppDB) {

}

func (o *AdapterKnockSensor433) OnSense(dev *api.Thing, service *api.ThingService, data api.ThingData) (state map[string]interface{}) {
	return nil
}

func (d *AdapterKnockSensor433) GetEventText(thing *api.Thing, sensor *api.ThingService, state map[string]interface{}) (shortText string, longText string) {
	return "", ""
}
