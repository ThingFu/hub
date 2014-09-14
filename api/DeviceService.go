// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Service for handling devices
type DeviceService interface {
	GetDevice(string) (Device, bool)
	SaveDevice(Device)
	GetDeviceType(string) DeviceType
	GetDeviceTypes() map[string]DeviceType
	RegisterDeviceType(DeviceType)
	RegisterDevice(Device)
	GetDevices() []Device
	Handle(*Device, *Sensor, map[string]interface {})

	Cycle()
	ContainerAware

	SetRulesService(RulesService)
	SetFactory(Factory)
	SetDataSource(DataSource)
}
