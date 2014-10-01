package tests

import "github.com/thingfu/hub/api"

type MockThingManager struct {
	thingTypes map[string]api.ThingType
	things     map[string]api.Thing

	rulesService api.RulesManager
	container    api.Container
	factory      api.Factory
	dataSource   api.DataSource
}

func (s *MockThingManager) SetDataSource(svc api.DataSource) {
	s.dataSource = svc
}

func (s *MockThingManager) SetRulesManager(svc api.RulesManager) {
	s.rulesService = svc
}

func (s *MockThingManager) SetFactory(o api.Factory) {
	s.factory = o
}

func (d *MockThingManager) GetContainer() api.Container {
	return d.container
}

func (d *MockThingManager) SetContainer(c api.Container) {
	d.container = c
}

func (o *MockThingManager) GetThing(id string) (dev api.Thing, ok bool) {
	dev, ok = o.things[id]

	return
}

func (o *MockThingManager) SaveThing(d api.Thing) {

}

func (o *MockThingManager) GetThingType(id string) api.ThingType {
	return o.thingTypes[id]
}

func (o *MockThingManager) GetThingTypes() map[string]api.ThingType {
	return o.thingTypes
}

func (o *MockThingManager) RegisterThingType(d api.ThingType) {

}

func (o *MockThingManager) RegisterThing(d api.Thing) {

}

func (o *MockThingManager) GetThings() []api.Thing {
	v := make(api.Things, len(o.things))

	return v
}

func (o *MockThingManager) Handle(thing *api.Thing, service *api.ThingService, state map[string]interface{}) {

}

func (o *MockThingManager) Cycle() {

}

func (o *MockThingManager) Actuate(t *api.Thing, op string, params map[string]interface{}) {

}

func (s *MockThingManager) ValidateWiring() {

}
