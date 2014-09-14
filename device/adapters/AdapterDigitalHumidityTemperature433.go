// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/go-home/hub/api"
	"log"
)

type AdapterDigitalHumidityTemperature433 struct {
}

func (d *AdapterDigitalHumidityTemperature433) Cycle(dev *api.Device) {

}

func (d *AdapterDigitalHumidityTemperature433) OnSense(dev *api.Device, data api.DeviceData) {
	log.Println("Sense WT450")

	// deviceManager := device.DeviceManagerInstance()
	/*
	   var $devMgr = $svcs.$devMgr;
	   var d = parseState(data);
	   $devMgr.updateState(dev.id, d);
	*/
}

func (d *AdapterDigitalHumidityTemperature433) GetEventText(*api.Device, *api.Sensor) (shortText string, longText string) {
	return "", ""
}

/*
func parseState (data RF433Data) map[string]float64 {
	return 0.0
}
*/

/*
function parseState (data) {
    var humidity = nutil.bin2dec(data.substring(13, 20));
    var temp1 = nutil.bin2dec(data.substring(20, 28)) - 50;
    var temp2 = nutil.bin2dec(data.substring(28, 35));

    return {
        humidity: humidity,
        tempMajor: temp1,
        tempMinor: temp2
    };
}

*/
