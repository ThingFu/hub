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
	api.BaseProtocolHandler
}

func (d *DelegatingSimulationProtocolHandler) OnStart() {

}

func (d *DelegatingSimulationProtocolHandler) OnStop() {

}

func (p *DelegatingSimulationProtocolHandler) GetName() string {
	return "sim"
}

func (p *DelegatingSimulationProtocolHandler) GetLabel() string {
	return "Simulation (For Debugging Only)"
}

func (p *DelegatingSimulationProtocolHandler) OnRead(payload api.ReadRequest) {

}

func (p *DelegatingSimulationProtocolHandler) Write(t *api.Thing, req api.WriteRequest) {

}
