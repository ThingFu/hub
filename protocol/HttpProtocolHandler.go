package protocol

import "github.com/go-home/hub/api"

type HttpProtocolHandler struct {
	factory       api.Factory
	environment   api.Environment
	thingService api.ThingService
	config        api.ProtocolConfiguration
}

func (p *HttpProtocolHandler) Start() {

}

func (p *HttpProtocolHandler) Stop() {

}

func (p *HttpProtocolHandler) SetFactory(o api.Factory) {
	p.factory = o
}

func (p *HttpProtocolHandler) SetThingService(o api.ThingService) {
	p.thingService = o
}

func (p *HttpProtocolHandler) SetEnvironment(o api.Environment) {
	p.environment = o
}

func (p *HttpProtocolHandler) SetProtocolConfiguration(o api.ProtocolConfiguration) {
	p.config = o
}

func (p *HttpProtocolHandler) IsEnabled() bool {
	return p.config.Enabled
}

func (p *HttpProtocolHandler) GetName() string {
	return "http"
}

func (p *HttpProtocolHandler) GetLabel() string {
	return "HTTP"
}

func (p *HttpProtocolHandler) Handle(payload interface{}) {

}

func (p *HttpProtocolHandler) SetContainer(api.Container) {

}

func (p *HttpProtocolHandler) ValidateWiring() {

}
