// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// DataSource for Devices, States and Events
type DataSource interface {
	GetDevices() []Device
	PutDevice(*Device)

	GetDeviceEventsCount() int
	GetDeviceEvents(int) []Event
	SaveDevice(Device)
	PutEvent(*Event)
	SaveState(*Device, map[string] interface {})
	ContainerAware

	SetEnvironment(Environment)
}
