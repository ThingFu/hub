package adapters

import (
	"fmt"
	"github.com/thingfu/hub/api"
	"log"
)

type Adapter4ButtonWireless433 struct {
}

func (d *Adapter4ButtonWireless433) OnWrite(t *api.Thing, op string, params api.WriteRequest, db api.AppDB, handler api.ProtocolHandler) {

}

func (d *Adapter4ButtonWireless433) Cycle(dev *api.Thing) {

}

func (d *Adapter4ButtonWireless433) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest, handler api.ProtocolHandler) (state api.State) {
	log.Println(fmt.Sprintf("4 Button Wireless Triggered %s", service.Label))

	return nil
}

func (d *Adapter4ButtonWireless433) GetEventText(dev *api.Thing, service *api.ThingService, state api.State) (shortText string, longText string) {
	shortText = fmt.Sprintf("Button %s on %s pressed", service.Label, dev.Name)
	longText = fmt.Sprintf("Button %s on %s pressed", service.Label, dev.Name)

	return
}
