// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"fmt"
	"github.com/thingfu/hub/api"
	"log"
)

type AdapterContactSensor433 struct {
}

func (s *AdapterContactSensor433) Cycle(dev *api.Thing) {

}

func (d *AdapterContactSensor433) OnWrite(t *api.Thing, op string, params api.WriteRequest, db api.AppDB, handler api.ProtocolHandler) {

}

func (sensor *AdapterContactSensor433) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest, handler api.ProtocolHandler) (state api.State) {
	log.Println("Sense : Contact Sensor")
	return nil
}

func (d *AdapterContactSensor433) GetEventText(dev *api.Thing, service *api.ThingService, state api.State) (shortText string, longText string) {
	txt := fmt.Sprintf("%s opened", dev.Name)
	shortText, longText = txt, txt

	return
}
