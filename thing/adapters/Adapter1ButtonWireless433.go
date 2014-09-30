// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"fmt"
	"github.com/thingfu/hub/api"
	"log"
)

type Adapter1ButtonWireless433 struct {
}

func (d *Adapter1ButtonWireless433) Cycle(dev *api.Thing) {

}

func (d *Adapter1ButtonWireless433) OnSense(dev *api.Thing, data api.ThingData) (state map[string]interface{}) {
	log.Println("Sense : Wireless Button")

	return nil
}

func (d *Adapter1ButtonWireless433) OnActuate(t *api.Thing, op string, params map[string]interface{}, db api.AppDB) {

}

func (d *Adapter1ButtonWireless433) GetEventText(dev *api.Thing, sensor *api.Sensor, state map[string]interface{}) (shortText string, longText string) {
	txt := fmt.Sprintf("Button %s pressed", dev.Name)
	shortText, longText = txt, txt

	return
}
