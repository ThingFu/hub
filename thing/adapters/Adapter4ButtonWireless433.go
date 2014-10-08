package adapters

import (
	"fmt"
	"github.com/thingfu/hub/api"
	"log"
)

type Adapter4ButtonWireless433 struct {
}

func (d *Adapter4ButtonWireless433) OnActuate(t *api.Thing, op string, params map[string]interface{}, db api.AppDB) {

}

func (d *Adapter4ButtonWireless433) Cycle(dev *api.Thing) {

}

func (d *Adapter4ButtonWireless433) OnSense(dev *api.Thing, service *api.ThingService, data api.ThingData) (state map[string]interface{}) {
	log.Println(fmt.Sprintf("4 Button Wireless Triggered %s", service.Label))

	return nil
}

func (d *Adapter4ButtonWireless433) GetEventText(dev *api.Thing, service *api.ThingService, state map[string]interface{}) (shortText string, longText string) {
	shortText = fmt.Sprintf("Button %s on %s pressed", service.Label, dev.Name)
	longText = fmt.Sprintf("Button %s on %s pressed", service.Label, dev.Name)

	return
}
