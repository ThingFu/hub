package tests

import "github.com/thingfu/hub/api"

type MockContainer struct {
	dataSource       api.DataSource
	rulesService     api.RulesManager
	thingManager     api.ThingManager
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

func (c *MockContainer) setRulesManager(o api.RulesManager) {
	c.rulesService = o
}

func (c *MockContainer) RulesManager() api.RulesManager {
	return c.rulesService
}

func (c *MockContainer) setThingManager(o api.ThingManager) {
	c.thingManager = o
}

func (c *MockContainer) ThingManager() api.ThingManager {
	return c.thingManager
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
