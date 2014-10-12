package api

type BaseCommunicationChannel struct {
	factory      Factory
	environment  Environment
	thingManager ThingManager
	config       ChannelConfiguration
	protocols	 map[string] ProtocolHandler
	container	 Container
}

func NewBaseCommunicationChannel() (BaseCommunicationChannel) {
	c := new(BaseCommunicationChannel)
	c.protocols = make(map[string] ProtocolHandler)

	return *c
}

func (cc BaseCommunicationChannel) GetProtocols() map[string] ProtocolHandler {
	return cc.protocols
}


func (cc BaseCommunicationChannel) GetConfiguration() ChannelConfiguration {
	return cc.config
}

func (cc *BaseCommunicationChannel) SetChannelConfiguration(c ChannelConfiguration) {
	cc.config = c
}

func (cc *BaseCommunicationChannel) SetThingManager(t ThingManager) {
	cc.thingManager = t
}

func (cc *BaseCommunicationChannel) SetFactory(f Factory) {
	cc.factory = f
}

func (cc *BaseCommunicationChannel) SetEnvironment(e Environment) {
	cc.environment = e
}

func (cc *BaseCommunicationChannel) SetContainer(c Container) {
	cc.container = c
}

func (cc *BaseCommunicationChannel) AddProtocol(h ProtocolHandler) {
	cc.protocols[h.GetName()] = h
}

func (s BaseCommunicationChannel) IsEnabled() bool   {
	return s.config.Enabled
}

