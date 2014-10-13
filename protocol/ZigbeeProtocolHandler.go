package protocol

import "github.com/thingfu/hub/api"

type ZigbeeProtocolHandler struct {
	api.BaseProtocolHandler
}

func (p *ZigbeeProtocolHandler) OnStart() {

}

func (p *ZigbeeProtocolHandler) OnStop() {

}

func (p *ZigbeeProtocolHandler) GetName() string {
	return "zigbee"
}

func (p *ZigbeeProtocolHandler) GetLabel() string {
	return "Zigbee"
}

func (p *ZigbeeProtocolHandler) OnRead(payload api.ReadRequest) {

}

func (p *ZigbeeProtocolHandler) Write(t *api.Thing, req api.WriteRequest) {

}

