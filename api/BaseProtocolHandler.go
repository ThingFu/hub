package api

type BaseProtocolHandler struct {
	factory      Factory
	environment  Environment
	thingManager ThingManager
}

func (h *BaseProtocolHandler) SetThingManager(t ThingManager) {
	h.thingManager = t
}

func (h *BaseProtocolHandler) SetFactory(f Factory) {
	h.factory = f
}

func (h *BaseProtocolHandler) SetEnvironment(e Environment) {
	h.environment = e
}

func (h *BaseProtocolHandler) GetThingManager() ThingManager{
	return h.thingManager
}

func (h *BaseProtocolHandler) GetFactory() Factory {
	return h.factory
}

func (h *BaseProtocolHandler) GetEnvironment() Environment {
	return h.environment
}

