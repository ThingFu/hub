// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/go-home/hub/api"
	"strconv"
)

type AdapterDigitalHumidityTemperature433 struct {

}

func (d *AdapterDigitalHumidityTemperature433) Cycle(dev *api.Device) {

}

func (d *AdapterDigitalHumidityTemperature433) OnSense(dev *api.Device, data api.DeviceData) (state map[string]interface{}) {
	dec, _ := strconv.Atoi(data.GetData()["dec"].(string))
	mask := 0x7f
	humidity := mask & (dec >> 16)

	mask = 0xff
	tempHigh := mask & (dec >> 8)
	tempLow := mask & dec

	state = make(map[string] interface {})

	state["h"] = humidity
	state["tH"] =  tempHigh-50
	state["tL"] = tempLow

	return
}

func (d *AdapterDigitalHumidityTemperature433) GetEventText(*api.Device, *api.Sensor) (shortText string, longText string) {
	return "", ""
}

