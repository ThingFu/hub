// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import _ "net/http"

// Device adapter to support any sensing or actuation
type DeviceAdapter interface {
	Cycle(*Device)
	OnSense(*Device, DeviceData) (state map[string]interface{})
	GetEventText(device *Device, sensor *Sensor, state map[string]interface{}) (shortText string, longText string)
	// HandleGet(req http.Request, res http.Response)
}
