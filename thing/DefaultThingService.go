// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package thing

import (
	"github.com/go-home/hub/api"
	"github.com/go-home/hub/utils"
	"log"
	"time"
	"sort"
)

type DefaultThingService struct {
	thingTypes map[string]api.ThingType
	things     map[string]api.Thing

	rulesService api.RulesService
	container    api.Container
	factory      api.Factory
	dataSource   api.DataSource
}

func NewThingService() *DefaultThingService {
	svc := new(DefaultThingService)
	svc.thingTypes = make(map[string]api.ThingType)
	svc.things = make(map[string]api.Thing)

	return svc
}

func (s *DefaultThingService) SetDataSource(svc api.DataSource) {
	s.dataSource = svc
}

func (s *DefaultThingService) SetRulesService(svc api.RulesService) {
	s.rulesService = svc
}

func (s *DefaultThingService) SetFactory(o api.Factory) {
	s.factory = o
}

func (d *DefaultThingService) GetContainer() api.Container {
	return d.container
}

func (d *DefaultThingService) SetContainer(c api.Container) {
	d.container = c
}

func (o *DefaultThingService) GetThing(id string) (dev api.Thing, ok bool) {
	dev, ok = o.things[id]

	return
}

func (o *DefaultThingService) SaveThing(d api.Thing) {
	go o.dataSource.SaveThing(d)
	o.things[d.GetId()] = d
}

func (o *DefaultThingService) GetThingType(id string) api.ThingType {
	return o.thingTypes[id]
}

func (o *DefaultThingService) GetThingTypes() map[string]api.ThingType {
	return o.thingTypes
}

func (o *DefaultThingService) RegisterThingType(d api.ThingType) {
	o.thingTypes[d.TypeId] = d
}

func (o *DefaultThingService) RegisterThing(d api.Thing) {
	descriptor := o.GetThingType(d.Type)
	d.Descriptor = descriptor
	o.things[d.GetId()] = d
}

func (o *DefaultThingService) GetThings() []api.Thing {
	v := make(api.Things, len(o.things))

	idx := 0
	for _, value := range o.things {
		v[idx] = value
		idx++
	}

	sort.Sort(v)

	return v
}

func (o *DefaultThingService) Handle(thing *api.Thing, sensor *api.Sensor, state map[string]interface{}) {
	facts := new(api.RuleFacts)
	facts.Thing = thing
	facts.Sensor = sensor
	facts.Target = thing.GetId()

	drv := o.factory.CreateThingAdapter(thing.Descriptor.TypeId)

	// Save latest state of thing
	if state != nil {
		state["lastUpdated"] = time.Now()
		o.dataSource.SaveState(thing, state)
	}

	// Save event to DB
	desc := thing.Descriptor
	if desc.LogEvents {
		evt := api.NewEvent()
		evt.Thing = thing.Id
		evt.Sensor = sensor.Name
		evt.ShortText, evt.LongText = drv.GetEventText(thing, sensor, state)
		evt.Event = api.EVENT_SENSE

		o.dataSource.PutEvent(evt)
	}

	// TODO: Publish Event to Cloud
	o.rulesService.Trigger(api.TRIGGER_THING, facts)
}

func (o *DefaultThingService) Cycle() {
	for _, dev := range o.things {
		thingType := dev.Type
		descriptor := dev.Descriptor
		cycletime := descriptor.CycleTime

		if cycletime > 0 {
			if utils.TimeWithinThreshold(dev.LastCycle, cycletime, 0) {
				drv := o.factory.CreateThingAdapter(thingType)
				if drv != nil {
					drv.Cycle(&dev)
				}
				dev.UpdateLastCycle(time.Now())
				o.SaveThing(dev)
			}
		}
	}
}

func (s *DefaultThingService) ValidateWiring() {
	if s.rulesService == nil {
		log.Fatal("[ERROR] rulesService not wired to DefaultThingService")
	}

	if s.dataSource == nil {
		log.Fatal("[ERROR] dataSource not wired to DefaultThingService")
	}

	if s.factory == nil {
		log.Fatal("[ERROR] factory not wired to DefaultThingService")
	}
}
