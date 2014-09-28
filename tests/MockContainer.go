package tests

import "github.com/thingfu/hub/api"

type MockContainer struct {
	dataSource       api.DataSource
	rulesService     api.RulesService
	thingService     api.ThingService
	scheduleService  api.ScheduleService
	environment      api.Environment
	factory          api.Factory
	protocolHandlers map[string]api.ProtocolHandler
}

func (c *MockContainer) setDataSource(o api.DataSource) {
	c.dataSource = o
}

func (c *MockContainer) DataSource() api.DataSource {
	return c.dataSource
}

func (c *MockContainer) setRulesService(o api.RulesService) {
	c.rulesService = o
}

func (c *MockContainer) RulesService() api.RulesService {
	return c.rulesService
}

func (c *MockContainer) setThingService(o api.ThingService) {
	c.thingService = o
}

func (c *MockContainer) ThingService() api.ThingService {
	return c.thingService
}

func (c *MockContainer) setEnvironment(o api.Environment) {
	c.environment = o
}

func (c *MockContainer) Env() api.Environment {
	return c.environment
}

func (c *MockContainer) setScheduleService(o api.ScheduleService) {
	c.scheduleService = o
}

func (c *MockContainer) ScheduleService() api.ScheduleService {
	return c.scheduleService
}

func (c *MockContainer) setFactory(o api.Factory) {
	c.factory = o
}

func (c *MockContainer) Factory() api.Factory {
	return c.factory
}

func (c *MockContainer) ProtocolHandlers() map[string]api.ProtocolHandler {
	return c.protocolHandlers
}

func (c *MockContainer) ProtocolHandler(p string) api.ProtocolHandler {
	return c.protocolHandlers[p]
}
