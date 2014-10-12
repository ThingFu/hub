package adapters

import (
	"github.com/thingfu/hub/api"
)

type AdapterThingFuDemoAdapter struct {

}

func (d *AdapterThingFuDemoAdapter) Cycle(dev *api.Thing) {

}

func (d *AdapterThingFuDemoAdapter) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest) (state map[string]interface{}) {
	return nil
}

func (d *AdapterThingFuDemoAdapter) OnWrite(t *api.Thing, op string, params api.WriteRequest, db api.AppDB) {

}

func (d *AdapterThingFuDemoAdapter) GetEventText(dev *api.Thing, service *api.ThingService, state map[string]interface{}) (shortText string, longText string) {
	return
}
