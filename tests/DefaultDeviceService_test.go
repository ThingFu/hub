package tests

/*
import (
	"fmt"
	"gohome/node/api"
	"gohome/node/device"
	"testing"
)

func TestRegisterDevices(t *testing.T) {
	ds := new(MockDataSource)
	env := new(MockEnvironment)

	deviceService := device.NewDefaultDeviceService(env, ds)

	deviceService.RegisterDevice(*api.NewDevice())
	deviceService.RegisterDevice(*api.NewDevice())

	ret := deviceService.GetDevices()
	if len(ret) != 1 {
		t.Fail()
	}
}
*/

/*
func NewDefaultDeviceService(env api.Environment, ds api.DataSource) *DefaultDeviceService {
	svc := new (DefaultDeviceService)
	svc.deviceTypes = make(map[string] api.DeviceType)
	svc.devices = make(map[string] api.Device)
	env.GetConfig()
	home := env.GetHome()
	registerDeviceTypes(home, svc)
	registerDevices(ds, svc);

	return svc
}
*/

/*
	fmt.Println("[INFO] Register Devices")
	devices := dataSource.GetDevices()

	for _, device := range devices {
		deviceService.RegisterDevice(device)
	}
}

*/
