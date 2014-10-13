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

func (d *AdapterKnockSensor433) OnWrite(t *api.Thing, op string, params api.WriteRequest, db api.AppDB, handler api.ProtocolHandler) {

}

func (o *AdapterKnockSensor433) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest, handler api.ProtocolHandler) (state api.State) {
	return nil
}

func (d *AdapterKnockSensor433) GetEventText(thing *api.Thing, sensor *api.ThingService, state api.State) (shortText string, longText string) {
	return "", ""
}
