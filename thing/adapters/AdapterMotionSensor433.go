// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"fmt"
	"github.com/go-home/hub/api"
)

type AdapterMotionSensor433 struct {
}

func (d *AdapterMotionSensor433) Cycle(dev *api.Thing) {

}

func (d *AdapterMotionSensor433) OnSense(dev *api.Thing, data api.ThingData) (state map[string]interface{}) {
	return nil
}

func (d *AdapterMotionSensor433) GetEventText(dev *api.Thing, sensor *api.Sensor, state map[string]interface{}) (shortText string, longText string) {
	txt := fmt.Sprintf("Motion by %s detected", dev.Name)
	shortText, longText = txt, txt

	return
}
