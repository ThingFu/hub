// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/go-home/hub/api"
)

type AdapterIPCamera struct {
}

func (d *AdapterIPCamera) Cycle(dev *api.Device) {

}

func (d *AdapterIPCamera) OnSense(dev *api.Device, data api.DeviceData) {

}

func (d *AdapterIPCamera) GetEventText(*api.Device, *api.Sensor) (shortText string, longText string) {
	return "", ""
}