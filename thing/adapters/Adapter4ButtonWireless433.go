package adapters

import (
	"fmt"
	"github.com/thingfu/hub/api"
)

type Adapter4ButtonWireless433 struct {
}

func (d *Adapter4ButtonWireless433) OnActuate(t *api.Thing, op string, params map[string]interface{}, db api.AppDB) {

}

func (d *Adapter4ButtonWireless433) Cycle(dev *api.Thing) {

}

func (d *Adapter4ButtonWireless433) OnSense(dev *api.Thing, data api.ThingData) (state map[string]interface{}) {
	return nil
}

func (d *Adapter4ButtonWireless433) GetEventText(dev *api.Thing, sensor *api.Sensor, state map[string]interface{}) (shortText string, longText string) {
	shortText = fmt.Sprintf("Button on %s pressed", dev.Name)
	longText = fmt.Sprintf("Button %s on %s pressed", sensor.Label, dev.Name)

	return
}
