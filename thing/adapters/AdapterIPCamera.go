// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/thingfu/hub/api"
)

type AdapterIPCamera struct {
}

func (d *AdapterIPCamera) OnWrite(t *api.Thing, op string, params api.WriteRequest, db api.AppDB, handler api.ProtocolHandler) {

}

func (d *AdapterIPCamera) Cycle(dev *api.Thing) {

}

func (d *AdapterIPCamera) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest, handler api.ProtocolHandler) (state api.State) {
	return nil
}

func (d *AdapterIPCamera) GetEventText(thing *api.Thing, sensor *api.ThingService, state api.State) (shortText string, longText string) {
	return "", ""
}
