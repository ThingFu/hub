package protocol

import "github.com/thingfu/hub/api"

type CodeMatchProtocolHandler struct {
}

func (cm *CodeMatchProtocolHandler) Start() {}
func (cm *CodeMatchProtocolHandler) Stop()  {}

func (cm *CodeMatchProtocolHandler) IsEnabled() bool {
	return true
}

func (cm *CodeMatchProtocolHandler) GetName() string {
	return "CodeMatch"
}

func (cm *CodeMatchProtocolHandler) GetLabel() string {
	return "Code Match"
}

func (cm *CodeMatchProtocolHandler) Handle(data interface{}) {

}

func (cm *CodeMatchProtocolHandler) SetProtocolConfiguration(api.ProtocolConfiguration) {}
func (cm *CodeMatchProtocolHandler) SetThingManager(api.ThingManager)                   {}
func (cm *CodeMatchProtocolHandler) SetFactory(api.Factory)                             {}
func (cm *CodeMatchProtocolHandler) SetEnvironment(api.Environment)                     {}
