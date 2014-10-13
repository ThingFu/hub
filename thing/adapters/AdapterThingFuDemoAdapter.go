package adapters

import (
	"github.com/thingfu/hub/api"
)

type AdapterThingFuDemoAdapter struct {

}

func (d *AdapterThingFuDemoAdapter) Cycle(dev *api.Thing) {

}

func (d *AdapterThingFuDemoAdapter) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest, handler api.ProtocolHandler) (state api.State) {
	return nil
}

func (d *AdapterThingFuDemoAdapter) OnWrite(t *api.Thing, action string, req api.WriteRequest, db api.AppDB, handler api.ProtocolHandler) {
	var content string
	switch action {
	case "turnOnRed":
		content = "111111100000000110100001"
		break

	case "turnOffRed":
		content = "111111100000000110100000"
		break

	case "turnOnYellow":
		content = "111111100000000110110001"
		break

	case "turnOffYellow":
		content = "111111100000000110110000"
		break
	}

	req.Put("content", content)

	handler.Write(t, req);
}

func (d *AdapterThingFuDemoAdapter) GetEventText(dev *api.Thing, service *api.ThingService, state api.State) (shortText string, longText string) {
	return
}
