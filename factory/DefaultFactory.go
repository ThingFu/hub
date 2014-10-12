// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/protocol"
	"github.com/thingfu/hub/rules/conditions"
	"github.com/thingfu/hub/rules/consequences"
	"github.com/thingfu/hub/thing/adapters"
	"github.com/thingfu/hub/channels"
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
	if t == "triggered" {
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

func (d *DefaultFactory) CreateChannelHandler(cfg api.ChannelConfiguration) (api.CommunicationChannel) {
	var channel api.CommunicationChannel
	switch cfg.Type {
	case "Serial":
		channel = channels.NewSerialChannel()
	}

	channel.SetThingManager(d.container.ThingManager())
	channel.SetChannelConfiguration(cfg)
	channel.SetFactory(d.container.Factory())

	return channel
}

func (d *DefaultFactory) CreateProtocolHandler(t string) (handler api.ProtocolHandler) {
	switch t {
	case "CodeMatch":
		handler = new(protocol.CodeMatchProtocolHandler)
		break;

	case "WT450":
		handler = new(protocol.WT450ProtocolHandler)
		break;

	case "ThingFuSimple":
		handler = new(protocol.ThingFuSimpleProtocolHandler)
		break;

	case "ZigBee":
		handler = new(protocol.ZigbeeProtocolHandler)
		break;

	case "ZWave":
		handler = new(protocol.ZWaveProtocolHandler)
		break;

	case "EnOcean":
		handler = new(protocol.EnOceanProtocolHandler)
		break;
	}

	handler.SetThingManager(d.container.ThingManager())
	handler.SetFactory(d.container.Factory())

	return
}

func (s *DefaultFactory) CreateThingAdapter(t string) api.ThingAdapter {
	var adapter api.ThingAdapter
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

	case "weather":
		adapter = new(adapters.AdapterWeather)

	case "utilities-webapp":
		adapter = new(adapters.AdapterWebAppUtilities)

	case "thingfudemo":
		adapter = new(adapters.AdapterThingFuDemoAdapter)

	default:
		return nil
	}
	return adapter
}

func (s *DefaultFactory) ValidateWiring() {

}
