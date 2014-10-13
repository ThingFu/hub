package protocol

import "github.com/thingfu/hub/api"

type EnOceanProtocolHandler struct {
	api.BaseProtocolHandler
}

func (p *EnOceanProtocolHandler) OnStart() {

}

func (p *EnOceanProtocolHandler) OnStop() {

}

func (p *EnOceanProtocolHandler) GetName() string {
	return "EnOcean"
}

func (p *EnOceanProtocolHandler) GetLabel() string {
	return "EnOcean"
}

func (p *EnOceanProtocolHandler) OnRead(payload api.ReadRequest) {

}

func (p *EnOceanProtocolHandler) Write(t *api.Thing, req api.WriteRequest) {

}

