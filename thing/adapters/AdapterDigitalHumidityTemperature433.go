// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"fmt"
	"github.com/thingfu/hub/api"
	"strconv"
)

type AdapterDigitalHumidityTemperature433 struct {
}

func (d *AdapterDigitalHumidityTemperature433) OnActuate(t *api.Thing, op string, params map[string]interface{}, db api.AppDB) {

}

func (d *AdapterDigitalHumidityTemperature433) Cycle(dev *api.Thing) {

}

func (d *AdapterDigitalHumidityTemperature433) OnSense(dev *api.Thing, service *api.ThingService, data api.ThingData) (state map[string]interface{}) {
	dec, _ := strconv.Atoi(data.GetData()["dec"].(string))
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

func (d *AdapterDigitalHumidityTemperature433) GetEventText(thing *api.Thing, service *api.ThingService, state map[string]interface{}) (shortText string, longText string) {
	humidity := state["h"]
	tH := state["tH"]
	tL := state["tL"]

	txt := fmt.Sprintf("Temperature %d.%d, Rel. Humidity %d%% @ %s", tH, tL, humidity, thing.Name)
	shortText, longText = txt, txt

	return
}
