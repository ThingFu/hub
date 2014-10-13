package api

type CommunicationChannel interface {
	Start() error
	IsEnabled() bool
	GetName() string
	GetLabel() string
	AddProtocol(ProtocolHandler)
	GetProtocol(string) (ProtocolHandler)
	GetProtocols() (map[string] ProtocolHandler)
	Write(WriteRequest)

	SetChannelConfiguration(ChannelConfiguration)
	SetThingManager(ThingManager)
	SetFactory(Factory)
	SetEnvironment(Environment)
	ContainerAware
}
