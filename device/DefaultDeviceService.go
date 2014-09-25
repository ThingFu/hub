// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"github.com/go-home/hub/api"
	"github.com/go-home/hub/utils"
	"log"
	"time"
)

type DefaultDeviceService struct {
	deviceTypes map[string]api.DeviceType
	devices     map[string]api.Device

	rulesService api.RulesService
	container    api.Container
	factory      api.Factory
	dataSource   api.DataSource
}

func NewDeviceService() *DefaultDeviceService {
	svc := new(DefaultDeviceService)
	svc.deviceTypes = make(map[string]api.DeviceType)
	svc.devices = make(map[string]api.Device)

	return svc
}

func (s *DefaultDeviceService) SetDataSource(svc api.DataSource) {
	s.dataSource = svc
}

func (s *DefaultDeviceService) SetRulesService(svc api.RulesService) {
	s.rulesService = svc
}

func (s *DefaultDeviceService) SetFactory(o api.Factory) {
	s.factory = o
}

func (d *DefaultDeviceService) GetContainer() api.Container {
	return d.container
}

func (d *DefaultDeviceService) SetContainer(c api.Container) {
	d.container = c
}

func (o *DefaultDeviceService) GetDevice(id string) (dev api.Device, ok bool) {
	dev, ok = o.devices[id]

	return
}

func (o *DefaultDeviceService) SaveDevice(d api.Device) {
	go o.dataSource.SaveDevice(d)
	o.devices[d.GetId()] = d
}

func (o *DefaultDeviceService) GetDeviceType(id string) api.DeviceType {
	return o.deviceTypes[id]
}

func (o *DefaultDeviceService) GetDeviceTypes() map[string]api.DeviceType {
	return o.deviceTypes
}

func (o *DefaultDeviceService) RegisterDeviceType(d api.DeviceType) {
	o.deviceTypes[d.TypeId] = d
}

func (o *DefaultDeviceService) RegisterDevice(d api.Device) {
	descriptor := o.GetDeviceType(d.Type)
	d.Descriptor = descriptor
	o.devices[d.GetId()] = d
}

func (o *DefaultDeviceService) GetDevices() []api.Device {
	v := make([]api.Device, len(o.devices))

	idx := 0
	for _, value := range o.devices {
		v[idx] = value
		idx++
	}
	return v
}

func (o *DefaultDeviceService) Handle(device *api.Device, sensor *api.Sensor, state map[string]interface{}) {
	facts := new(api.RuleFacts)
	facts.Device = device
	facts.Sensor = sensor
	facts.Target = device.GetId()

	drv := o.factory.CreateDeviceAdapter(device.Descriptor.TypeId)

	// Save latest state of device
	if state != nil {
		state["lastUpdated"] = time.Now()
		o.dataSource.SaveState(device, state)
	}

	// Save event to DB
	desc := device.Descriptor
	if desc.LogEvents {
		evt := api.NewEvent()
		evt.Device = device.Id
		evt.Sensor = sensor.Name
		evt.ShortText, evt.LongText = drv.GetEventText(device, sensor, state)
		evt.Event = api.EVENT_SENSE

		o.dataSource.PutEvent(evt)
	}

	// TODO: Publish Event to Cloud
	o.rulesService.Trigger(api.TRIGGER_DEVICE, facts)
}

func (o *DefaultDeviceService) Cycle() {
	for _, dev := range o.devices {
		deviceType := dev.Type
		descriptor := dev.Descriptor
		cycletime := descriptor.CycleTime

		if cycletime > 0 {
			if utils.TimeWithinThreshold(dev.LastCycle, cycletime, 0) {
				drv := o.factory.CreateDeviceAdapter(deviceType)
				if drv != nil {
					drv.Cycle(&dev)
				}
				dev.UpdateLastCycle(time.Now())
				o.SaveDevice(dev)
			}
		}
	}
}

func (s *DefaultDeviceService) ValidateWiring() {
	if s.rulesService == nil {
		log.Fatal("[ERROR] rulesService not wired to DefaultDeviceService")
	}

	if s.dataSource == nil {
		log.Fatal("[ERROR] dataSource not wired to DefaultDeviceService")
	}

	if s.factory == nil {
		log.Fatal("[ERROR] factory not wired to DefaultDeviceService")
	}
}
