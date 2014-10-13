package protocol

import "github.com/thingfu/hub/api"

type ZWaveProtocolHandler struct {
	api.BaseProtocolHandler
}

func (p *ZWaveProtocolHandler) OnStart() {

}

func (p *ZWaveProtocolHandler) OnStop() {

}

func (p *ZWaveProtocolHandler) GetName() string {
	return "ZWave"
}

func (p *ZWaveProtocolHandler) GetLabel() string {
	return "ZWave"
}

func (p *ZWaveProtocolHandler) OnRead(payload api.ReadRequest) {

}

func (p *ZWaveProtocolHandler) Write(t *api.Thing, req api.WriteRequest) {

}

