package adapters

import (
	"fmt"
	"github.com/go-home/hub/api"
)

type Adapter4ButtonWireless433 struct {
}

func (d *Adapter4ButtonWireless433) Cycle(dev *api.Device) {

}

func (d *Adapter4ButtonWireless433) OnSense(dev *api.Device, data api.DeviceData) {

}

func (d *Adapter4ButtonWireless433) GetEventText(dev *api.Device, sensor *api.Sensor) (shortText string, longText string) {
	shortText = fmt.Sprintf("Button on %s pressed", dev.Name)
	longText = fmt.Sprintf("Button %s on %s pressed", sensor.Label, dev.Name)

	return
}
