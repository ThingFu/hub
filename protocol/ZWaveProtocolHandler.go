package protocol

import "github.com/thingfu/hub/api"

type ZWaveProtocolHandler struct {
	factory      api.Factory
	environment  api.Environment
	thingManager api.ThingManager
	config       api.ProtocolConfiguration
}

func (p *ZWaveProtocolHandler) IsEnabled() bool {
	return p.config.Enabled
}

func (p *ZWaveProtocolHandler) Start() {

}

func (p *ZWaveProtocolHandler) Stop() {

}

func (p *ZWaveProtocolHandler) SetFactory(o api.Factory) {
	p.factory = o
}

func (p *ZWaveProtocolHandler) SetThingManager(o api.ThingManager) {
	p.thingManager = o
}

func (p *ZWaveProtocolHandler) SetEnvironment(o api.Environment) {
	p.environment = o
}

func (p *ZWaveProtocolHandler) SetProtocolConfiguration(o api.ProtocolConfiguration) {
	p.config = o
}

func (p *ZWaveProtocolHandler) GetName() string {
	return "ZWave"
}

func (p *ZWaveProtocolHandler) GetLabel() string {
	return "ZWave"
}

func (p *ZWaveProtocolHandler) Handle(payload interface{}) {

}

func (p *ZWaveProtocolHandler) SetContainer(api.Container) {

}

func (p *ZWaveProtocolHandler) ValidateWiring() {

}
