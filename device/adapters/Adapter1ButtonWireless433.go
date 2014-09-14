// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/go-home/hub/api"
	"log"
	"fmt"
)

type Adapter1ButtonWireless433 struct {
}

func (d *Adapter1ButtonWireless433) Cycle(dev *api.Device) {

}

func (d *Adapter1ButtonWireless433) OnSense(dev *api.Device, data api.DeviceData) {
	log.Println("Sense : Wireless Button")
}

func (d *Adapter1ButtonWireless433) GetEventText(dev *api.Device, *api.Sensor) (shortText string, longText string) {
	shortText, longText = fmt.Sprintf("Button %s pressed", dev.Name)

	return
}


