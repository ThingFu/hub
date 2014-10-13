// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package container

import (
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/data/source"
	"github.com/thingfu/hub/env"
	"github.com/thingfu/hub/events"
	"github.com/thingfu/hub/factory"
	"github.com/thingfu/hub/rules"
	"github.com/thingfu/hub/thing"
	"log"
	"github.com/thingfu/hub/channels"
)

var CONTAINER *DefaultContainer = nil

func Instance() api.Container {
	return CONTAINER
}

func Initialize(home string, config api.Configuration) (api.Container, api.Environment) {
	CONTAINER = new(DefaultContainer)

	env := env.NewEnvironment(home, config)

	CONTAINER.Register(env, "api.Environment")
	CONTAINER.Register(rules.NewRulesManager(), "api.RulesManager")
	CONTAINER.Register(thing.NewThingManager(), "api.ThingManager")
	CONTAINER.Register(new(events.DefaultScheduleService), "api.ScheduleService")
	CONTAINER.Register(source.NewMongoDataSource(), "api.DataSource")
	CONTAINER.Register(new(factory.DefaultFactory), "api.Factory")
	CONTAINER.Register(channels.NewCommChannelManager(), "api.CommChannelManager")

	CONTAINER.startWire()

	return CONTAINER, env
}

type DefaultContainer struct {
	dataSource       api.DataSource
	rulesService     api.RulesManager
	thingManager     api.ThingManager
	scheduleService  api.ScheduleService
	environment      api.Environment
	factory          api.Factory
	commManager		 api.CommChannelManager
	protocolHandlers map[string] api.ProtocolHandler
	channels 		 map[string] api.CommunicationChannel
}

func (c *DefaultContainer) Register(svc api.ContainerAware, t string) {
	switch {
	case t == "api.Environment":
		c.environment = svc.(api.Environment)

	case t == "api.RulesManager":
		c.rulesService = svc.(api.RulesManager)

	case t == "api.ThingManager":
		c.thingManager = svc.(api.ThingManager)

	case t == "api.ScheduleService":
		c.scheduleService = svc.(api.ScheduleService)

	case t == "api.DataSource":
		c.dataSource = svc.(api.DataSource)

	case t == "api.Factory":
		c.factory = svc.(api.Factory)

	case t == "api.ProtocolHandler":
		name := svc.(api.ProtocolHandler).GetName()
		c.protocolHandlers[name] = svc.(api.ProtocolHandler)

	case t == "api.CommChannelManager":
		c.commManager = svc.(api.CommChannelManager)

	default:
		log.Println("Unknown Service")
	}
}

func (c *DefaultContainer) startWire() {
	rulesService := c.RulesManager()
	factory := c.Factory()
	thingManager := c.ThingManager()
	env := c.Env()
	dataSource := c.DataSource()
	scheduleServices := c.ScheduleService()
	commManager := c.CommChannelManager()

	// Wire Up Services
	// Rules Service
	rulesService.SetThingManager(thingManager)
	rulesService.SetFactory(factory)

	// Comm Channel Manager
	commManager.SetFactory(factory)

	// Factory

	// ThingManager
	thingManager.SetRulesManager(rulesService)
	thingManager.SetFactory(factory)
	thingManager.SetDataSource(dataSource)
	thingManager.SetCommChannelManager(commManager)

	// DataSource
	dataSource.SetEnvironment(env)

	// ScheduleService
	scheduleServices.SetRulesManager(rulesService)
	scheduleServices.SetThingManager(thingManager)

	// Protocol Handlers
	c.protocolHandlers = make(map[string]api.ProtocolHandler)
	c.channels = make(map[string] api.CommunicationChannel)

	services := make([]api.ContainerAware, 6)
	services[0] = rulesService
	services[1] = factory
	services[2] = thingManager
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

func (c *DefaultContainer) setRulesManager(o api.RulesManager) {
	c.rulesService = o
}

func (c *DefaultContainer) RulesManager() api.RulesManager {
	return c.rulesService
}

func (c *DefaultContainer) setThingManager(o api.ThingManager) {
	c.thingManager = o
}

func (c *DefaultContainer) ThingManager() api.ThingManager {
	return c.thingManager
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

func (c *DefaultContainer) CommChannelManager() api.CommChannelManager {
	return c.commManager
}

func (c *DefaultContainer) ProtocolHandlers() map[string]api.ProtocolHandler {
	return c.protocolHandlers
}

func (c *DefaultContainer) ProtocolHandler(p string) api.ProtocolHandler {
	return c.protocolHandlers[p]
}

func (c *DefaultContainer) Channels() map[string]api.CommunicationChannel {
	return c.channels
}

func (c *DefaultContainer) Channel(s string) api.CommunicationChannel {
	return c.channels[s]
}


