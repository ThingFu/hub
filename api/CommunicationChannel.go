package api

type CommunicationChannel interface {
	Start() error
	IsEnabled() bool
	GetName() string
	GetLabel() string
	AddProtocol(ProtocolHandler)

	SetChannelConfiguration(ChannelConfiguration)
	SetThingManager(ThingManager)
	SetFactory(Factory)
	SetEnvironment(Environment)
	ContainerAware
}
