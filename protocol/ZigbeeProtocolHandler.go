package protocol

import "github.com/go-home/hub/api"

type ZigbeeProtocolHandler struct {
	factory       api.Factory
	environment   api.Environment
	deviceService api.DeviceService
	config        api.ProtocolConfiguration
}

func (p *ZigbeeProtocolHandler) IsEnabled() bool {
	return p.config.Enabled
}

func (p *ZigbeeProtocolHandler) Start() {

}

func (p *ZigbeeProtocolHandler) Stop() {

}

func (p *ZigbeeProtocolHandler) SetFactory(o api.Factory) {
	p.factory = o
}

func (p *ZigbeeProtocolHandler) SetDeviceService(o api.DeviceService) {
	p.deviceService = o
}

func (p *ZigbeeProtocolHandler) SetEnvironment(o api.Environment) {
	p.environment = o
}

func (p *ZigbeeProtocolHandler) SetProtocolConfiguration(o api.ProtocolConfiguration) {
	p.config = o
}

func (p *ZigbeeProtocolHandler) GetName() string {
	return "zigbee"
}

func (p *ZigbeeProtocolHandler) GetLabel() string {
	return "Zigbee"
}

func (p *ZigbeeProtocolHandler) Handle(payload interface{}) {

}

func (p *ZigbeeProtocolHandler) SetContainer(api.Container) {

}

func (p *ZigbeeProtocolHandler) ValidateWiring() {

}
