// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"log"
	"github.com/thingfu/hub/api"
)

type AdapterWebAppUtilities struct {
}

func (d *AdapterWebAppUtilities) OnWrite(t *api.Thing, op string, req api.WriteRequest, db api.AppDB, handler api.ProtocolHandler) {
	db.Put(op, req)

	log.Println("AdapterWebAppUtilities -- Actuate")
}

func (d *AdapterWebAppUtilities) Cycle(dev *api.Thing) {
	log.Println("AdapterWebAppUtilities -- Cycle")
}

func (d *AdapterWebAppUtilities) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest, handler api.ProtocolHandler) (state api.State) {
	return nil
}

func (d *AdapterWebAppUtilities) GetEventText(dev *api.Thing, sensor *api.ThingService, state api.State) (shortText string, longText string) {
	return
}
