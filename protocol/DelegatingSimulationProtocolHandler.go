package protocol

import (
	"github.com/thingfu/hub/api"
)

// Protocol used by the Simulation URL
// This constructs and delegates to actual Protocol Handlers
// This is meant for debugging purposes. Not a good idea
// to expose this live
//
type DelegatingSimulationProtocolHandler struct {
	factory      api.Factory
	environment  api.Environment
	thingManager api.ThingManager
	config       api.ProtocolConfiguration
}

func (d *DelegatingSimulationProtocolHandler) Start() {

}

func (d *DelegatingSimulationProtocolHandler) Stop() {

}

func (d *DelegatingSimulationProtocolHandler) SetProtocolConfiguration(config api.ProtocolConfiguration) {
	d.config = config
}

func (d *DelegatingSimulationProtocolHandler) SetThingManager(thingManager api.ThingManager) {
	d.thingManager = thingManager
}

func (d *DelegatingSimulationProtocolHandler) SetFactory(factory api.Factory) {
	d.factory = factory
}

func (d *DelegatingSimulationProtocolHandler) SetEnvironment(env api.Environment) {
	d.environment = env
}

func (p *DelegatingSimulationProtocolHandler) IsEnabled() bool {
	return p.config.Enabled
}

func (p *DelegatingSimulationProtocolHandler) GetName() string {
	return "sim"
}

func (p *DelegatingSimulationProtocolHandler) GetLabel() string {
	return "Simulation (For Debugging Only)"
}

func (p *DelegatingSimulationProtocolHandler) Handle(payload interface{}) {

}
func (p *DelegatingSimulationProtocolHandler) SetContainer(api.Container) {

}

func (p *DelegatingSimulationProtocolHandler) ValidateWiring() {

}
