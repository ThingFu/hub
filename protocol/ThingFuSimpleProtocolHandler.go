package protocol

import (
	"log"
	"github.com/thingfu/hub/api"
	"errors"
)

type ThingFuSimpleProtocolHandler struct {
	api.BaseProtocolHandler
}

func (p *ThingFuSimpleProtocolHandler) OnStop() {

}

func (p *ThingFuSimpleProtocolHandler) OnStart() {
	log.Println("[INFO] Starting..")
}

func (p *ThingFuSimpleProtocolHandler) OnRead(payload api.ReadRequest) {

}

func (p *ThingFuSimpleProtocolHandler) getThing(ser string) (*api.Thing, *api.ThingService, error) {
	return new(api.Thing), new(api.ThingService), errors.New("Unknown Thing")
}

func (p *ThingFuSimpleProtocolHandler) GetName() string {
	return "ThingFuSimple"
}

func (p *ThingFuSimpleProtocolHandler) GetLabel() string {
	return "ThingFu Simple"
}

