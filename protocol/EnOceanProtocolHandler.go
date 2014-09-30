package protocol

import "github.com/thingfu/hub/api"

type EnOceanProtocolHandler struct {
	factory      api.Factory
	environment  api.Environment
	thingManager api.ThingManager
	config       api.ProtocolConfiguration
}

func (p *EnOceanProtocolHandler) IsEnabled() bool {
	return p.config.Enabled
}

func (p *EnOceanProtocolHandler) Start() {

}

func (p *EnOceanProtocolHandler) Stop() {

}

func (p *EnOceanProtocolHandler) SetFactory(o api.Factory) {
	p.factory = o
}

func (p *EnOceanProtocolHandler) SetThingManager(o api.ThingManager) {
	p.thingManager = o
}

func (p *EnOceanProtocolHandler) SetEnvironment(o api.Environment) {
	p.environment = o
}

func (p *EnOceanProtocolHandler) SetProtocolConfiguration(o api.ProtocolConfiguration) {
	p.config = o
}

func (p *EnOceanProtocolHandler) GetName() string {
	return "EnOcean"
}

func (p *EnOceanProtocolHandler) GetLabel() string {
	return "EnOcean"
}

func (p *EnOceanProtocolHandler) Handle(payload interface{}) {

}

func (p *EnOceanProtocolHandler) SetContainer(api.Container) {

}

func (p *EnOceanProtocolHandler) ValidateWiring() {

}
