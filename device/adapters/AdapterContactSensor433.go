// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import "github.com/go-home/hub/api"

type AdapterContactSensor433 struct {
}

func (s *AdapterContactSensor433) Cycle(dev *api.Device) {

}

func (sensor *AdapterContactSensor433) OnSense(dev *api.Device, data api.DeviceData) {

}

func (d *AdapterContactSensor433) GetEventText(*api.Device, *api.Sensor) (shortText string, longText string) {
	return "", ""
}