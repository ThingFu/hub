// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package thing

import (
	"fmt"
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/utils"
	"log"
	"sort"
	"time"
)

type DefaultThingManager struct {
	thingTypes map[string]api.ThingType
	things     map[string]api.Thing

	rulesService api.RulesManager
	container    api.Container
	factory      api.Factory
	dataSource   api.DataSource
	commManager	 api.CommChannelManager
}

func NewThingManager() *DefaultThingManager {
	svc := new(DefaultThingManager)
	svc.thingTypes = make(map[string]api.ThingType)
	svc.things = make(map[string]api.Thing)

	return svc
}

func (s *DefaultThingManager) SetCommChannelManager(svc api.CommChannelManager) {
	s.commManager = svc
}


func (s *DefaultThingManager) SetDataSource(svc api.DataSource) {
	s.dataSource = svc
}

func (s *DefaultThingManager) SetRulesManager(svc api.RulesManager) {
	s.rulesService = svc
}

func (s *DefaultThingManager) SetFactory(o api.Factory) {
	s.factory = o
}

func (d *DefaultThingManager) GetContainer() api.Container {
	return d.container
}

func (d *DefaultThingManager) SetContainer(c api.Container) {
	d.container = c
}

func (o *DefaultThingManager) GetThing(id string) (dev api.Thing, ok bool) {
	dev, ok = o.things[id]

	return
}

func (o *DefaultThingManager) SaveThing(d api.Thing) {
	go o.dataSource.SaveThing(d)
	o.things[d.GetId()] = d
}

func (o *DefaultThingManager) RemoveThing(t api.Thing) {
	o.dataSource.DeleteThing(t)
	delete(o.things, t.Id)
}

func (o *DefaultThingManager) GetThingType(id string) (t api.ThingType, e error) {
	t = o.thingTypes[id]

	if &t == nil {
		e = fmt.Errorf("Unknown type %s", id)
	}
	return
}

func (o *DefaultThingManager) GetThingTypes() map[string]api.ThingType {
	return o.thingTypes
}

func (o *DefaultThingManager) RegisterThingType(d api.ThingType) {
	o.thingTypes[d.TypeId] = d
}

func (o *DefaultThingManager) RegisterThing(d api.Thing) {
	descriptor, err := o.GetThingType(d.Type)
	if err != nil {
		log.Println(err)
	}

	d.Descriptor = descriptor
	o.things[d.GetId()] = d
}

func (o *DefaultThingManager) GetThings() []api.Thing {
	v := make(api.Things, len(o.things))

	idx := 0
	for _, value := range o.things {
		v[idx] = value
		idx++
	}

	sort.Sort(v)

	return v
}

func (o *DefaultThingManager) CreateThing(t *api.Thing) {
	n := o.dataSource.PutThing(t)
	n.Descriptor, _ = o.GetThingType(t.Type)
	o.things[n.Id] = n
}

func (o *DefaultThingManager) LoadThings() {
	things := o.dataSource.GetThings()
	for _, thing := range things {
		o.RegisterThing(thing)
	}
}

func (o *DefaultThingManager) Handle(thing *api.Thing, service *api.ThingService, state api.State) {
	facts := new(api.RuleFacts)
	facts.Thing = thing
	facts.Service = service
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
		evt.Service = service.Name
		evt.ShortText, evt.LongText = drv.GetEventText(thing, service, state)
		evt.Event = api.EVENT_SENSE
		evt.Data = state

		o.dataSource.PutEvent(evt)
	}

	// Run Rules for this thing
	o.rulesService.Trigger(api.TRIGGER_THING, facts)
}

func (o *DefaultThingManager) Actuate(t *api.Thing, action string, params map[string]interface{}) {
	adapter := o.factory.CreateThingAdapter(t.Type)

	appDB := o.dataSource.CreateAppDB(t)

	req := api.NewWriteRequest(params)
	req.Put("action", action)

	for k, v := range params {
		req.Put(k, v)
	}

	handler := o.GetProtocolHandlerForThing(t)
	adapter.OnWrite(t, action, req, appDB, handler)
}

func (o *DefaultThingManager) Cycle() {
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

func (o *DefaultThingManager) GetProtocolHandlerForThing(t *api.Thing) (api.ProtocolHandler) {
	protocol := t.Descriptor.Protocol

	ch := o.commManager.GetChannelForProtocol(protocol)

	return ch.GetProtocol(protocol)
}

func (s *DefaultThingManager) ValidateWiring() {
	if s.rulesService == nil {
		log.Fatal("[ERROR] rulesService not wired to DefaultThingManager")
	}

	if s.dataSource == nil {
		log.Fatal("[ERROR] dataSource not wired to DefaultThingManager")
	}

	if s.factory == nil {
		log.Fatal("[ERROR] factory not wired to DefaultThingManager")
	}
}
