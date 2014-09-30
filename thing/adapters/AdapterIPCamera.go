// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/thingfu/hub/api"
)

type AdapterIPCamera struct {
}

func (d *AdapterIPCamera) OnActuate(t *api.Thing, op string, params map[string]interface{}, db api.AppDB) {

}

func (d *AdapterIPCamera) Cycle(dev *api.Thing) {

}

func (d *AdapterIPCamera) OnSense(dev *api.Thing, data api.ThingData) (state map[string]interface{}) {
	return nil
}

func (d *AdapterIPCamera) GetEventText(thing *api.Thing, sensor *api.Sensor, state map[string]interface{}) (shortText string, longText string) {
	return "", ""
}
