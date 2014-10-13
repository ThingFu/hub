// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"fmt"
	"github.com/thingfu/hub/api"
)

type AdapterDigitalHumidityTemperature433 struct {
}

func (d *AdapterDigitalHumidityTemperature433) OnWrite(t *api.Thing, op string, params api.WriteRequest, db api.AppDB, handler api.ProtocolHandler) {

}

func (d *AdapterDigitalHumidityTemperature433) Cycle(dev *api.Thing) {

}

func (d *AdapterDigitalHumidityTemperature433) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest, handler api.ProtocolHandler) (state api.State) {
	fmt.Println(data.GetPayload())

	dec := data.GetAsInt("dhtdata")
	mask := 0x7f
	humidity := mask & (dec >> 16)

	mask = 0xff
	tempHigh := mask & (dec >> 8)
	tempLow := mask & dec

	state = make(map[string]interface{})

	state["h"] = humidity
	state["tH"] = tempHigh - 50
	state["tL"] = tempLow

	return
}

func (d *AdapterDigitalHumidityTemperature433) GetEventText(thing *api.Thing, service *api.ThingService, state api.State) (shortText string, longText string) {
	humidity := state["h"]
	tH := state["tH"]
	tL := state["tL"]

	txt := fmt.Sprintf("Temperature %d.%d, Rel. Humidity %d%% @ %s", tH, tL, humidity, thing.Name)
	shortText, longText = txt, txt

	return
}
