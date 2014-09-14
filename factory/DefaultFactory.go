// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"github.com/go-home/hub/api"
	"github.com/go-home/hub/device/adapters"
	"github.com/go-home/hub/protocol"
	"github.com/go-home/hub/rules/conditions"
	"github.com/go-home/hub/rules/consequences"
)

type DefaultFactory struct {
	container api.Container
}

func (d *DefaultFactory) GetContainer() api.Container {
	return d.container
}

func (d *DefaultFactory) SetContainer(c api.Container) {
	d.container = c
}

func (d *DefaultFactory) CreateCondition(t string) api.Condition {
	var condition api.Condition
	if t == "sense" {
		condition = new(conditions.Sense)
	} else if t == "hourly" {
		condition = new(conditions.Hourly)
	}

	return condition
}

func (d *DefaultFactory) CreateConsequence(t string) api.Consequence {
	var consequence api.Consequence
	switch t {
	case "sendmail":
		consequence = new(consequences.SendMail)

	case "logwrite":
		consequence = new(consequences.LogWrite)
	}
	return consequence
}

func (d *DefaultFactory) CreateProtocolHandler(t string, cfg api.ProtocolConfiguration) api.ProtocolHandler {
	var handler api.ProtocolHandler
	switch t {
	case "RF433":
		handler = new(protocol.RF433ProtocolHandler)

	case "http":
		handler = new(protocol.HttpProtocolHandler)

	case "zigbee":
		handler = new(protocol.ZigbeeProtocolHandler)

	case "sim":
		handler = new(protocol.DelegatingSimulationProtocolHandler)
	}
	handler.SetDeviceService(d.container.DeviceService())
	handler.SetProtocolConfiguration(cfg)
	handler.SetFactory(d.container.Factory())

	return handler
}

func (s *DefaultFactory) CreateDeviceAdapter(t string) api.DeviceAdapter {
	var adapter api.DeviceAdapter
	switch t {
	case "lgtv-47ls5700":
		adapter = new(adapters.AdapterLGTV47LS5700)

	case "433mhz-contact":
		adapter = new(adapters.AdapterContactSensor433)

	case "433mhz-1button":
		adapter = new(adapters.Adapter1ButtonWireless433)

	case "433mhz-wt450":
		adapter = new(adapters.AdapterDigitalHumidityTemperature433)

	case "433mhz-knock":
		adapter = new(adapters.AdapterKnockSensor433)

	case "433mhz-motion":
		adapter = new(adapters.AdapterMotionSensor433)

	case "dlink-dcs930l":
		adapter = new(adapters.AdapterIPCamera)

	case "433mhz-4buttons":
		adapter = new(adapters.Adapter4ButtonWireless433)

	default:
		return nil
	}
	return adapter
}

func (s *DefaultFactory) ValidateWiring() {

}
