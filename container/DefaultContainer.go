// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package container

import (
	"github.com/go-home/hub/api"
	"github.com/go-home/hub/data/source"
	"github.com/go-home/hub/device"
	"github.com/go-home/hub/env"
	"github.com/go-home/hub/events"
	"github.com/go-home/hub/factory"
	"github.com/go-home/hub/rules"
	"log"
)

var CONTAINER *DefaultContainer = nil

func Instance() api.Container {
	return CONTAINER
}

func Initialize(home string, config api.Configuration) (api.Container, api.Environment) {
	CONTAINER = new(DefaultContainer)

	env := env.NewEnvironment(home, config)
	CONTAINER.Register(env, "api.Environment")
	CONTAINER.Register(rules.NewRulesService(), "api.RulesService")
	CONTAINER.Register(device.NewDeviceService(), "api.DeviceService")
	CONTAINER.Register(new(events.DefaultScheduleService), "api.ScheduleService")
	CONTAINER.Register(new(source.MongoDataSource), "api.DataSource")
	CONTAINER.Register(new(factory.DefaultFactory), "api.Factory")

	CONTAINER.startWire()

	// Register Protocol Handlers
	protocols := config.Protocols
	for k, protocol := range protocols {
		handler := CONTAINER.Factory().CreateProtocolHandler(k, protocol)
		CONTAINER.Register(handler.(api.ContainerAware), "api.ProtocolHandler")
	}

	return CONTAINER, env
}

type DefaultContainer struct {
	dataSource       api.DataSource
	rulesService     api.RulesService
	deviceService    api.DeviceService
	scheduleService  api.ScheduleService
	environment      api.Environment
	factory          api.Factory
	protocolHandlers map[string]api.ProtocolHandler
}

func (c *DefaultContainer) Register(svc api.ContainerAware, t string) {
	switch {
	case t == "api.Environment":
		c.environment = svc.(api.Environment)

	case t == "api.RulesService":
		c.rulesService = svc.(api.RulesService)

	case t == "api.DeviceService":
		c.deviceService = svc.(api.DeviceService)

	case t == "api.ScheduleService":
		c.scheduleService = svc.(api.ScheduleService)

	case t == "api.DataSource":
		c.dataSource = svc.(api.DataSource)

	case t == "api.Factory":
		c.factory = svc.(api.Factory)

	case t == "api.ProtocolHandler":
		name := svc.(api.ProtocolHandler).GetName()
		c.protocolHandlers[name] = svc.(api.ProtocolHandler)

	default:
		log.Println("Unknown Service")
	}
}

func (c *DefaultContainer) startWire() {
	rulesService := c.RulesService()
	factory := c.Factory()
	deviceService := c.DeviceService()
	env := c.Env()
	dataSource := c.DataSource()
	scheduleServices := c.ScheduleService()

	// Wire Up Services
	// Rules Service
	rulesService.SetDeviceService(deviceService)
	rulesService.SetFactory(factory)

	// Factory

	// DeviceService
	deviceService.SetRulesService(rulesService)
	deviceService.SetFactory(factory)
	deviceService.SetDataSource(dataSource)

	// DataSource
	dataSource.SetEnvironment(env)

	// ScheduleService
	scheduleServices.SetRulesService(rulesService)
	scheduleServices.SetDeviceService(deviceService)

	// Protocol Handlers
	c.protocolHandlers = make(map[string]api.ProtocolHandler)

	services := make([]api.ContainerAware, 6)
	services[0] = rulesService
	services[1] = factory
	services[2] = deviceService
	services[3] = env
	services[4] = dataSource
	services[5] = scheduleServices

	// Inject Container into all ContainerAware Services
	for _, service := range services {
		service.SetContainer(c)
	}

	// Validate Wiring
	for _, service := range services {
		service.ValidateWiring()
	}
}

func (c *DefaultContainer) setDataSource(o api.DataSource) {
	c.dataSource = o
}

func (c *DefaultContainer) DataSource() api.DataSource {
	return c.dataSource
}

func (c *DefaultContainer) setRulesService(o api.RulesService) {
	c.rulesService = o
}

func (c *DefaultContainer) RulesService() api.RulesService {
	return c.rulesService
}

func (c *DefaultContainer) setDeviceService(o api.DeviceService) {
	c.deviceService = o
}

func (c *DefaultContainer) DeviceService() api.DeviceService {
	return c.deviceService
}

func (c *DefaultContainer) setEnvironment(o api.Environment) {
	c.environment = o
}

func (c *DefaultContainer) Env() api.Environment {
	return c.environment
}

func (c *DefaultContainer) setScheduleService(o api.ScheduleService) {
	c.scheduleService = o
}

func (c *DefaultContainer) ScheduleService() api.ScheduleService {
	return c.scheduleService
}

func (c *DefaultContainer) setFactory(o api.Factory) {
	c.factory = o
}

func (c *DefaultContainer) Factory() api.Factory {
	return c.factory
}

func (c *DefaultContainer) ProtocolHandlers() map[string]api.ProtocolHandler {
	return c.protocolHandlers
}

func (c *DefaultContainer) ProtocolHandler(p string) api.ProtocolHandler {
	return c.protocolHandlers[p]
}
