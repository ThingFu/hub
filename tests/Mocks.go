package tests

import (
	"gohome/node/api"
)

type MockDataSource struct {

}

func (e *MockDataSource) GetDevices() []api.Device {
	return make([]api.Device, 0)
}

func (e *MockDataSource) PutDevice(dev *api.Device) {

}

func (e *MockDataSource) GetDeviceEvents(limit int) []api.Event {
	return make([]api.Event, 0)
}

func (e *MockDataSource) SaveDevice(device api.Device) {

}

func (e *MockDataSource) PutEvent() {

}

type MockEnvironment struct {

}


func (e *MockEnvironment) GetUptime() string {
	return ""
}

func (e *MockEnvironment) GetConfig() api.Configuration {
	cfg := new (api.Configuration)
	cfg.ServerPort = 8181

	return *cfg
}

func (e *MockEnvironment) GetHome() string {
	return ""
}

func (e *MockEnvironment) IsStartedUp() bool {
	return true
}
