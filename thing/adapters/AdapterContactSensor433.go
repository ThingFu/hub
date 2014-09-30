// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"fmt"
	"github.com/thingfu/hub/api"
)

type AdapterContactSensor433 struct {
}

func (s *AdapterContactSensor433) Cycle(dev *api.Thing) {

}

func (d *AdapterContactSensor433) OnActuate(t *api.Thing, op string, params map[string]interface{}, db api.AppDB) {

}

func (sensor *AdapterContactSensor433) OnSense(dev *api.Thing, data api.ThingData) (state map[string]interface{}) {
	return nil
}

func (d *AdapterContactSensor433) GetEventText(dev *api.Thing, service *api.ThingService, state map[string]interface{}) (shortText string, longText string) {
	txt := fmt.Sprintf("%s opened", dev.Name)
	shortText, longText = txt, txt

	return
}
