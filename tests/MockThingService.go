package tests

import "github.com/thingfu/hub/api"

type MockThingService struct {
	thingTypes map[string]api.ThingType
	things     map[string]api.Thing

	rulesService api.RulesService
	container    api.Container
	factory      api.Factory
	dataSource   api.DataSource
}

func (s *MockThingService) SetDataSource(svc api.DataSource) {
	s.dataSource = svc
}

func (s *MockThingService) SetRulesService(svc api.RulesService) {
	s.rulesService = svc
}

func (s *MockThingService) SetFactory(o api.Factory) {
	s.factory = o
}

func (d *MockThingService) GetContainer() api.Container {
	return d.container
}

func (d *MockThingService) SetContainer(c api.Container) {
	d.container = c
}

func (o *MockThingService) GetThing(id string) (dev api.Thing, ok bool) {
	dev, ok = o.things[id]

	return
}

func (o *MockThingService) SaveThing(d api.Thing) {

}

func (o *MockThingService) GetThingType(id string) api.ThingType {
	return o.thingTypes[id]
}

func (o *MockThingService) GetThingTypes() map[string]api.ThingType {
	return o.thingTypes
}

func (o *MockThingService) RegisterThingType(d api.ThingType) {

}

func (o *MockThingService) RegisterThing(d api.Thing) {

}

func (o *MockThingService) GetThings() []api.Thing {
	v := make(api.Things, len(o.things))

	return v
}

func (o *MockThingService) Handle(thing *api.Thing, sensor *api.Sensor, state map[string]interface{}) {

}

func (o *MockThingService) Cycle() {

}

func (s *MockThingService) ValidateWiring() {

}
