// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/go-home/hub/api"
)

type AdapterKnockSensor433 struct {
}

func (o *AdapterKnockSensor433) Cycle(dev *api.Device) {

}

func (o *AdapterKnockSensor433) OnSense(dev *api.Device, data api.DeviceData) {

}

func (d *AdapterKnockSensor433) GetEventText(*api.Device, *api.Sensor) (shortText string, longText string) {
	return "", ""
}
