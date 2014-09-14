// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/go-home/hub/api"
)

type AdapterMotionSensor433 struct {
}

func (d *AdapterMotionSensor433) Cycle(dev *api.Device) {

}

func (d *AdapterMotionSensor433) OnSense(dev *api.Device, data api.DeviceData) {
	/*
	   var $devMgr = $svcs.$devMgr;
	   var $db = $svcs.$db;

	   dev.icon =  "imoon imoon-feed text-blue";

	   $db.putEvent(dev, "sense", "Motion Detected.");
	   $devMgr.updateState(dev.id, {});

	*/
}

func (d *AdapterMotionSensor433) GetEventText(*api.Device, *api.Sensor) (shortText string, longText string) {
	return "", ""
}
